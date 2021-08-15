package weatherapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	Sunny                              = 1000
	PartlyCloudy                       = 1003
	Cloudy                             = 1006
	Overcast                           = 1009
	Mist                               = 1030
	PatchyRainPossible                 = 1063
	PatchySnowPossible                 = 1066
	PatchySleetPossible                = 1069
	PatchyFreezingDrizzlePossible      = 1072
	ThunderyOutbreaksPossible          = 1087
	BlowingSnow                        = 1114
	Blizzard                           = 1117
	Fog                                = 1135
	FreezingFog                        = 1147
	PatchyLightDrizzle                 = 1150
	LightDrizzle                       = 1153
	FreezingDrizzle                    = 1168
	HeavyFreezingDrizzle               = 1171
	PatchyLightRain                    = 1180
	LightRain                          = 1183
	ModerateRainAtTimes                = 1186
	ModerateRain                       = 1189
	HeavyRainAtTimes                   = 1192
	HeavyRain                          = 1195
	LightFreezingRain                  = 1198
	ModerateOrHeavyFreezingRain        = 1201
	LightSleet                         = 1204
	ModerateOrHeavySleet               = 1207
	PatchyLightSnow                    = 1210
	LightSnow                          = 1213
	PatchyModerateSnow                 = 1216
	ModerateSnow                       = 1219
	PatchyHeavySnow                    = 1222
	HeavySnow                          = 1225
	IcePellets                         = 1237
	LightRainShower                    = 1240
	ModerateOrHeavyRainShower          = 1243
	TorrentialRainShower               = 1246
	LightSleetShowers                  = 1249
	ModerateOrHeavySleetShowers        = 1252
	LightSnowShowers                   = 1255
	ModerateOrHeavySnowShowers         = 1258
	LightShowersOfIcePellets           = 1261
	ModerateOrHeavyShowersOfIcePellets = 1264
	PatchyLightRainWithThunder         = 1273
	ModerateOrHeavyRainWithThunder     = 1276
	PatchyLightSnowWithThunder         = 1279
	ModerateOrHeavySnowWithThunder     = 1282
)

type Condition struct {
	ConditionText string `json:"text"`
	ConditionCode int    `json:"code"`
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
	ChanceOfRain           int       `json:"daily_chance_of_rain"`
	Condition              Condition `json:"condition"`
}

type ForecastDay struct {
	Date string `json:"date"`
	Day  Day    `json:"day"`
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
