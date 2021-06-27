package presentation

import (
	"fmt"
	"weatherapiCLI/apis/weatherapi"
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

func isCloudy(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.Cloudy, weatherapi.PartlyCloudy, weatherapi.Overcast:
		return true
	}
	return false
}

func isFog(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.Fog, weatherapi.Mist, weatherapi.FreezingFog:
		return true
	}
	return false
}

func isLightRain(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.LightRain, weatherapi.PatchyLightRain,
		weatherapi.PatchyRainPossible, weatherapi.LightFreezingRain,
		weatherapi.LightRainShower, weatherapi.PatchyLightDrizzle,
		weatherapi.LightDrizzle, weatherapi.FreezingDrizzle,
		weatherapi.HeavyFreezingDrizzle, weatherapi.ModerateOrHeavyRainShower,
		weatherapi.PatchyFreezingDrizzlePossible:
		return true
	}
	return false
}

func isHeavyRain(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.HeavyRain, weatherapi.HeavyRainAtTimes, weatherapi.ModerateRain,
		weatherapi.ModerateRainAtTimes, weatherapi.TorrentialRainShower,
		weatherapi.ModerateOrHeavyFreezingRain:
		return true
	}
	return false
}

func isLightSnow(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.LightSnow, weatherapi.LightSnowShowers,
		weatherapi.PatchyLightSnow, weatherapi.PatchySnowPossible,
		weatherapi.LightShowersOfIcePellets, weatherapi.BlowingSnow,
		weatherapi.IcePellets:
		return true
	}
	return false
}

func isHeavySnow(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.HeavySnow, weatherapi.ModerateOrHeavySnowShowers,
		weatherapi.PatchyModerateSnow, weatherapi.ModerateSnow,
		weatherapi.PatchyHeavySnow, weatherapi.ModerateOrHeavyShowersOfIcePellets,
		weatherapi.Blizzard:
		return true
	}
	return false
}

func isSleet(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.LightSleet, weatherapi.LightSleetShowers,
		weatherapi.ModerateOrHeavySleetShowers, weatherapi.PatchySleetPossible,
		weatherapi.ModerateOrHeavySleet:
		return true
	}
	return false
}

func isSunny(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.Sunny:
		return true
	}
	return false
}

func isHeavyRainThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.ModerateOrHeavyRainWithThunder, weatherapi.ThunderyOutbreaksPossible:
		return true
	}
	return false
}

func isHeavyShowersThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.PatchyLightRainWithThunder:
		return true
	}
	return false
}

func isSnowShowerThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case weatherapi.PatchyLightSnowWithThunder, weatherapi.ModerateOrHeavySnowWithThunder:
		return true
	}
	return false
}

func getWeatherAsciiDrawing(weatherConditionCode int) []string {
	switch {
	case isSunny(weatherConditionCode):
		return sunnyAscii
	case isCloudy(weatherConditionCode):
		return cloudyAscii
	case isFog(weatherConditionCode):
		return fogAscii
	case isLightRain(weatherConditionCode):
		return lightRainAscii
	case isHeavyRain(weatherConditionCode):
		return heavyRainAscii
	case isLightSnow(weatherConditionCode):
		return lightSnowAscii
	case isHeavySnow(weatherConditionCode):
		return heavySnowAscii
	case isSleet(weatherConditionCode):
		return sleetAscii
	case isHeavyRainThunder(weatherConditionCode):
		return thunderyHeavyRainAscii
	case isHeavyShowersThunder(weatherConditionCode):
		return thunderyHeavyShowersAscii
	case isSnowShowerThunder(weatherConditionCode):
		return thunderySnowShowersAscii
	default:
		return cloudyAscii
	}
}
