package weatherapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Condition struct {
	ConditionText string `json:"text"`
}

type CurrentWeather struct {
	LastUpdated         string    `json:"last_updated"`
	Temperature_Celsius float32   `json:"temp_c"`
	Wind_kph            float32   `json:"wind_kph"`
	Humidity            int       `json:"humidity"`
	Condition           Condition `json:"condition"`
}

type Day struct {
	MaxTemperature_Celsius float32   `json:"maxtemp_c"`
	MinTemperature_Celsius float32   `json:"mintemp_c"`
	AverageHumidity        float32   `json:"avghumidity"`
	ChanceOfRain           string       `json:"daily_chance_of_rain"`
	Condition              Condition `json:"condition"`
}

type Hour struct {
	DateAndTime         string    `json:"time"`
	Temperature_Celsius float32   `json:"temp_c"`
	Humidity            int       `json:"humidity"`
	Clouds              int       `json:"cloud"`
	ChanceOfRain        string       `json:"chance_of_rain"`
	Condition           Condition `json:"condition"`
}

type ForecastDay struct {
	Date  string `json:"date"`
	Day   Day    `json:"day"`
	Hours []Hour `json:"hour"`
}

type Forecast struct {
	ForecastDays []ForecastDay `json:"forecastday"`
}

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type WeatherApiData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
	Forecast Forecast       `json:"forecast"`
}

func (apiClient *WeatherClient) GetForecast(ctx context.Context, options *ForecastOptions) (*WeatherApiData, error) {
	numberOfDays := 1
	location := "auto:ip"
	if options != nil {
		numberOfDays = options.NumberOfDays
		location = options.Location
	}
	request, err := http.NewRequest("GET",
		fmt.Sprintf("%s/forecast.json", apiClient.BaseURL), nil)
	parameters := url.Values{}
	parameters.Add("key", apiClient.apiKey)
	parameters.Add("q", location)
	parameters.Add("days", fmt.Sprint(numberOfDays))
	parameters.Add("lang", options.Language)
	parameters.Add("aqi", "no")
	parameters.Add("alerts", "no")
	request.URL.RawQuery = parameters.Encode()
	if err != nil {
		return nil, err
	}

	return apiClient.sendRequest(ctx, request)
}

func (apiClient *WeatherClient) GetCurrent(ctx context.Context, options *ForecastOptions) (*WeatherApiData, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/current.json", apiClient.BaseURL), nil)

	parameters := url.Values{}
	parameters.Add("key", apiClient.apiKey)
	parameters.Add("q", options.Location)
	parameters.Add("lang", options.Language)
	parameters.Add("aqi", "no")
	request.URL.RawQuery = parameters.Encode()

	if err != nil {
		return nil, err
	}
	return apiClient.sendRequest(ctx, request)
}
