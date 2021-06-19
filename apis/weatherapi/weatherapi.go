package weatherapi

import (
	"encoding/json"
	"net/http"
	"os"
)

type Condition struct {
	ConditionText string `json:"text"`
}

type CurrentWeather struct {
	DateAndTime         string    `json:"last_updated"`
	Temperature_Celsius float32   `json:"temp_c"`
	Wind_kph            float32   `json:"wind_kph"`
	Humidity            int       `json:"humidity"`
	Condition           Condition `json:"condition"`
}

type Day struct {
	MaxTemperature_Celsius float32   `json:"maxtemp_c"`
	MinTemperature_Celsius float32   `json:"mintemp_c"`
	AverageHumidity        float32   `json:"avghumidity"`
	ChanceOfRain           int       `json:"daily_chance_of_raind"`
	Condition              Condition `json:"condition"`
}

type Hour struct {
	DateAndTime         string    `json:"time"`
	Temperature_Celsius float32   `json:"temp_c"`
	Humidity            int       `json:"humidity"`
	Clouds              int       `json:"cloud"`
	ChanceOfRain        int       `json:"daily_chance_of_raind"`
	Condition           Condition `json:"condition"`
}


type ForecastDay struct {
	Date  string `json:"date"`
	Day   Day    `json:"day"`
	Hours []Hour `json:"hour"`
}

type Forecast struct{
    ForecastDays []ForecastDay `json:"forecastday"`
}

type Location struct{
    Name string `json:"name"`
    Region string `json:"region"`
}

type WeatherApiData struct {
	Currrent CurrentWeather `json:"current"`
    Forecast Forecast `json:"forecast"`
}

func getWeatherData(location string, wantForecast bool, numberOfDays int32) (WeatherApiData, error) {
	const baseURL = "https://api.weatherapi.com/v1/"
	client := http.Client{}
	requestType := "current.json"
	if wantForecast == true {
		requestType = "forecast.json"
	}
	request, err := http.NewRequest("GET", baseURL+requestType, nil)
	if err != nil {
		return WeatherApiData{}, err
	}
	apiKey := os.Getenv("MY_WEATHER_API_KEY")
	query := request.URL.Query()
	query.Add("key", apiKey)
	query.Add("q", location)
	query.Add("aqi", "no")

	if wantForecast == true {
		query.Add("days", string(numberOfDays)) //nubmerOfDays has to be between 1 and 10
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		return WeatherApiData{}, err
	}
	var weatherData WeatherApiData
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&weatherData)
	if err != nil {
		return WeatherApiData{}, err
	}
	return weatherData, nil
}

func GetCurrentWeather(location string) (WeatherApiData, error) {
	return getWeatherData(location, false, 0)
}

func GetForecast(location string, numberOfDays int32) (WeatherApiData, error) {
	return getWeatherData(location, true, numberOfDays)
}
