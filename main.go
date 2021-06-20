package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"weatherapiCLI/apis/weatherapi"
	"weatherapiCLI/config"
)

type CommandLineOptions struct {
	NumberOfDays int
	Location     string
	Forecast     bool
}

func parseCommandLineArguments() *CommandLineOptions {
	numberOfDays := flag.Int("days", 1, "Number of days for the forecast")
	location := flag.String("location", "auto:ip", "Location for weather data(default is location via ip)")
	forecast := flag.Bool("forecast", false, "Boolean that represents a request for forecast")
	flag.Parse()
	return &CommandLineOptions{
		NumberOfDays: *numberOfDays,
		Location:     *location,
		Forecast:     *forecast,
	}
}

func printForecast(apiData *weatherapi.WeatherApiData) {
	printCurrentWeather(apiData)
	for i := 0; i < 24; i += 3 {
		fmt.Printf("At %s condition is: %s with a temperature of: %.1f\n",
			apiData.Forecast.ForecastDays[0].Hours[i].DateAndTime,
			apiData.Forecast.ForecastDays[0].Hours[i].Condition.ConditionText,
			apiData.Forecast.ForecastDays[0].Hours[i].Temperature_Celsius)
	}
}

func printCurrentWeather(apiData *weatherapi.WeatherApiData) {
	fmt.Printf("In %s, %s, %s\n", apiData.Location.Name, apiData.Location.Region, apiData.Location.Country)
	fmt.Printf("%s with temperature of %.1f Celsius\n", apiData.Currrent.Condition.ConditionText, apiData.Currrent.Temperature_Celsius)
}

func main() {
	configuration, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	options := parseCommandLineArguments()
	if options.Location == "auto:ip" {
		options.Location = configuration.Location
	}

	client := weatherapi.NewWeatherClient(os.Getenv("MY_WEATHER_API_KEY"))
	if options.Forecast == true {
		data, err := client.GetForecast(context.Background(), &weatherapi.ForecastOptions{NumberOfDays: options.NumberOfDays, Location: options.Location})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printForecast(data)
	} else {
		data, err := client.GetCurrent(context.Background(), options.Location)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printCurrentWeather(data)
	}
}
