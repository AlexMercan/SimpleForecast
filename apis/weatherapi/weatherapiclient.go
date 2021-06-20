package weatherapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

const (
	BaseURLV1 = "https://api.weatherapi.com/v1"
)

type WeatherClient struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type ForecastOptions struct {
	NumberOfDays int
	Location     string
}

func NewWeatherClient(key string) *WeatherClient {
	return &WeatherClient{
		BaseURL:    BaseURLV1,
		apiKey:     key,
		HTTPClient: &http.Client{},
	}
}

func (apiClient *WeatherClient) sendRequest(ctx context.Context, request *http.Request) (*WeatherApiData, error) {

	request = request.WithContext(ctx)
	request.Header.Set("Accept", "application/json")
	response, err := apiClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return nil, fmt.Errorf("Unknown error, status code %d", response.StatusCode)
		}
		return nil, errors.New(errorResponse.Error.Message)
	}
	var responseData WeatherApiData
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return nil, err
	}
	return &responseData, nil
}
