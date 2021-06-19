package main

import (
	"fmt"
	"os"
	"weatherapiCLI/apis/weatherapi"
)

func main() {
	data, err := weatherapi.GetForecast("tulcea", 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("In %s, %s, %s\n", data.Location.Name, data.Location.Region, data.Location.Country)
	fmt.Printf("%s with temperature of %.1f Celsius\n", data.Currrent.Condition.ConditionText, data.Currrent.Temperature_Celsius)
	fmt.Printf("At %s condition is: %s with a temperature of: %.1f\n",
		data.Forecast.ForecastDays[0].Hours[19].DateAndTime,
		data.Forecast.ForecastDays[0].Hours[19].Condition.ConditionText,
		data.Forecast.ForecastDays[0].Hours[19].Temperature_Celsius)
}
