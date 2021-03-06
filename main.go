package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"simpleforecast/apis/weatherapi"
	"simpleforecast/config"
	"simpleforecast/presentation"
)

type CommandLineOptions struct {
	NumberOfDays int
	Location     string
	Forecast     bool
	Language     string
}

func parseCommandLineArguments(config *config.Config) *CommandLineOptions {
	numberOfDays := flag.Int("days", 0, "Number of days for the forecast(1 to 3)")
	location := flag.String("location", config.Location, "Location for weather data(default is location via ip)")
	lang := flag.String("lang", config.Language, "Language of the weather status(sunny, cloudy, etc.)")
	flag.Parse()
	return &CommandLineOptions{
		NumberOfDays: *numberOfDays,
		Location:     *location,
		Forecast:     (*numberOfDays > 0),
		Language:     *lang,
	}
}

func main() {
	configuration, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	options := parseCommandLineArguments(configuration)

	client := weatherapi.NewWeatherClient(os.Getenv("MY_WEATHER_API_KEY"), weatherapi.BaseURLV1)
	if options.Forecast == true {
		data, err := client.GetForecast(context.TODO(),
			&weatherapi.ForecastOptions{NumberOfDays: options.NumberOfDays, Location: options.Location, Language: options.Language})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		presentation.PrintForecast(data)
	} else {
		data, err := client.GetCurrent(context.TODO(),
			&weatherapi.ForecastOptions{Location: options.Location, Language: options.Language})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		presentation.PrintCurrentWeather(data)
	}
}
