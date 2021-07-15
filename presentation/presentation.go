package presentation

import (
	"fmt"
	"simpleForecast/apis/weatherapi"
)

func PrintForecast(apiData *weatherapi.WeatherApiData) {
	fmt.Printf("In %s, %s, %s\n", apiData.Location.Name, apiData.Location.Region, apiData.Location.Country)
	for i := 0; i < len(apiData.Forecast.ForecastDays); i++ {
		fmt.Printf("Date: %s\n", apiData.Forecast.ForecastDays[i].Date)
		printForecastDay(&apiData.Forecast.ForecastDays[i])
		fmt.Println("")
	}
}

func printForecastDay(forecastDayData *weatherapi.ForecastDay) {
	coloredTemperature := colorTemperature(forecastDayData.Day.MaxTemperature_Celsius)
	asciiDrawing := getWeatherAsciiDrawing(forecastDayData.Day.Condition.ConditionCode)
	fmt.Printf("%sMax Temperature: %s°C\n", asciiDrawing[0], coloredTemperature)
	coloredTemperature = colorTemperature(forecastDayData.Day.MinTemperature_Celsius)
	fmt.Printf("%sMin Temperature: %s°C\n", asciiDrawing[1], coloredTemperature)
	coloredHumidity := applyColor(forecastDayData.Day.AverageHumidity, Blue)
	fmt.Printf("%sHumidity: %s%%\n", asciiDrawing[2], coloredHumidity)
	fmt.Printf("%sDaily chance of rain: %s%%\n", asciiDrawing[3], ColorStr(forecastDayData.Day.ChanceOfRain, Blue))
	fmt.Printf("%sStatus: %s\n", asciiDrawing[4], forecastDayData.Day.Condition.ConditionText)
}

func PrintCurrentWeather(apiData *weatherapi.WeatherApiData) {
	asciiDrawing := getWeatherAsciiDrawing(apiData.Current.Condition.ConditionCode)
	fmt.Printf("%sIn %s, %s, %s\n", asciiDrawing[0], apiData.Location.Name, apiData.Location.Region, apiData.Location.Country)
	coloredTemperature := colorTemperature(apiData.Current.Temperature_Celsius)
	fmt.Printf("%sTemperature: %s°C\n", asciiDrawing[1], coloredTemperature)
	coloredHumidity := applyColor(apiData.Current.Humidity, Blue)
	fmt.Printf("%sHumidity: %s%%\n", asciiDrawing[2], coloredHumidity)
	fmt.Printf("%sWind: %s kph \n", asciiDrawing[3], applyColor(apiData.Current.Wind_kph, Green))
	fmt.Printf("%sStatus: %s\n", asciiDrawing[4], apiData.Current.Condition.ConditionText)
}

//Applies(and returns) the color "colorCode" to the default string formatting of v
func applyColor(v interface{}, colorCode string) string {
	return ColorStr(fmt.Sprintf("%v", v), colorCode)
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

func getWeatherAsciiDrawing(weatherConditionCode int) []string {
	switch {
	case weatherapi.IsSunny(weatherConditionCode):
		return sunnyAscii
	case weatherapi.IsCloudy(weatherConditionCode):
		return cloudyAscii
	case weatherapi.IsFog(weatherConditionCode):
		return fogAscii
	case weatherapi.IsLightRain(weatherConditionCode):
		return lightRainAscii
	case weatherapi.IsHeavyRain(weatherConditionCode):
		return heavyRainAscii
	case weatherapi.IsLightSnow(weatherConditionCode):
		return lightSnowAscii
	case weatherapi.IsHeavySnow(weatherConditionCode):
		return heavySnowAscii
	case weatherapi.IsSleet(weatherConditionCode):
		return sleetAscii
	case weatherapi.IsHeavyRainThunder(weatherConditionCode):
		return thunderyHeavyRainAscii
	case weatherapi.IsHeavyShowersThunder(weatherConditionCode):
		return thunderyHeavyShowersAscii
	case weatherapi.IsSnowShowerThunder(weatherConditionCode):
		return thunderySnowShowersAscii
	default:
		return cloudyAscii
	}
}
