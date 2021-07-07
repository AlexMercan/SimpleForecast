package weatherapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGetCurrentWeather(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("./testData/currentWeatherTestData.json")
		if err != nil {
			panic(err)
		}
		w.Write(data)
	}))
	defer server.Close()
	apiClient := NewWeatherClient(os.Getenv("MY_WEATHER_API_KEY"), server.URL)
	data, err := apiClient.GetCurrent(context.Background(), &ForecastOptions{Location: "New York", Language: "en"})
	if err != nil {
		t.Errorf("Got unexpected error from current weather request: %s\n", err.Error())
	}
	fileData, err := os.ReadFile("./testData/currentWeatherTestData.json")
	var expectedData WeatherApiData
	err = json.Unmarshal(fileData, &expectedData)
	if err != nil {
		t.Errorf("Error while parsing expected test data file\n")
	}
	if !reflect.DeepEqual(*data, expectedData) {
		t.Errorf("Data != expected data\n")
	}
}

func TestGetForecast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("./testData/forecastWeatherTestData.json")
		if err != nil {
			panic(err)
		}
		w.Write(data)
	}))
	defer server.Close()
	apiClient := NewWeatherClient(os.Getenv("MY_WEATHER_API_KEY"), server.URL)
	data, err := apiClient.GetForecast(context.Background(),
		&ForecastOptions{NumberOfDays: 3, Location: "New York", Language: "en"})
	if err != nil {
		t.Errorf("Got unexpected error from forecast weather request\n")
	}
	fileData, err := os.ReadFile("./testData/forecastWeatherTestData.json")
	var expectedData WeatherApiData
	err = json.Unmarshal(fileData, &expectedData)
	if err != nil {
		t.Errorf("Error while parsing expected test data file\n")
	}
	if !reflect.DeepEqual(*data, expectedData) {
		t.Errorf("Data !=expectedData\n")
	}
}
