package presentation

import (
	"fmt"
	"weatherapiCLI/apis/weatherapi"
)

func PrintForecast(apiData *weatherapi.WeatherApiData) {
	fmt.Printf("In %s, %s, %s\n", apiData.Location.Name, apiData.Location.Region, apiData.Location.Country)
    for i:=0;i<len(apiData.Forecast.ForecastDays); i++{
        fmt.Printf("Date: %s\n", apiData.Forecast.ForecastDays[i].Date)
        printForecastDay(&apiData.Forecast.ForecastDays[i])
        fmt.Println("")
    }
}

func printForecastDay(forecastDayData *weatherapi.ForecastDay) {
    coloredTemperature := colorTemperature(forecastDayData.Day.MaxTemperature_Celsius)
    fmt.Printf("Max Temperature: %s°C\n", coloredTemperature)
    coloredTemperature = colorTemperature(forecastDayData.Day.MinTemperature_Celsius)
    fmt.Printf("Min Temperature: %s°C\n", coloredTemperature)
	humidityStr := fmt.Sprintf("%.1f%%", forecastDayData.Day.AverageHumidity)
	coloredHumidity := ColorStr(humidityStr, Blue)
    fmt.Printf("Humidity: %s%%\n", coloredHumidity)
    fmt.Printf("Daily chance of rain: %s%%\n", ColorStr(forecastDayData.Day.ChanceOfRain, Blue))
	fmt.Printf("Status: %s\n", forecastDayData.Day.Condition.ConditionText)
}

func PrintCurrentWeather(apiData *weatherapi.WeatherApiData) {
	fmt.Printf("In %s, %s, %s\n", apiData.Location.Name, apiData.Location.Region, apiData.Location.Country)
	coloredTemperature := colorTemperature(apiData.Current.Temperature_Celsius)
	fmt.Printf("Temperature: %s°C\n", coloredTemperature)
	humidityStr := fmt.Sprintf("%d%%", apiData.Current.Humidity)
	coloredHumidity := ColorStr(humidityStr, Blue)
	fmt.Printf("Humidity: %s \n", coloredHumidity)
	fmt.Printf("Status: %s\n", apiData.Current.Condition.ConditionText)
}

//Returns a colored string that corresponds to a temperature.
//If the temperature is above 14.0 degrees, it will be colored in red, otherwise it will be colored in blue
func colorTemperature(temperature float32) string {
	temperatureStr := fmt.Sprintf("%.1f", temperature)
	coloredTemperature := ColorStr(temperatureStr, Red)
	if temperature < 14.0 {
		coloredTemperature = ColorStr(temperatureStr, Blue)
	}
	return coloredTemperature
}
