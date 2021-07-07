package presentation

import (
	"reflect"
	"testing"
	"weatherapiCLI/apis/weatherapi"
)

func TestGetAsciiDrawing(t *testing.T) {
	cases := []struct {
		input    int
		expected []string
	}{
		{weatherapi.PartlyCloudy, cloudyAscii},
		{weatherapi.Fog, fogAscii},
		{weatherapi.LightRain, lightRainAscii},
		{weatherapi.LightRainShower, lightRainAscii},
		{weatherapi.HeavyRain, heavyRainAscii},
		{weatherapi.HeavyRainAtTimes, heavyRainAscii},
		{weatherapi.HeavySnow, heavySnowAscii},
		{weatherapi.ModerateSnow, heavySnowAscii},
		{weatherapi.Blizzard, heavySnowAscii},
		{weatherapi.LightSnow, lightSnowAscii},
		{weatherapi.LightSnowShowers, lightSnowAscii},
		{weatherapi.LightSleet, sleetAscii},
		{weatherapi.LightSleetShowers, sleetAscii}, //I'll change this to test
		{weatherapi.Sunny, sunnyAscii},
		{weatherapi.ModerateOrHeavyRainWithThunder, thunderyHeavyRainAscii},
		{weatherapi.ModerateOrHeavySnowWithThunder, thunderySnowShowersAscii},
		{weatherapi.PatchyLightRainWithThunder, thunderyHeavyShowersAscii},
	}
	for _, testCase := range cases {
		actual := getWeatherAsciiDrawing(testCase.input)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Expected: \n")
			for _, line := range testCase.expected {
				t.Errorf("%s", line)
			}
			t.Errorf("Got: \n")
			for _, line := range actual {
				t.Errorf("%s", line)
			}
			t.Errorf("For: %d\n", testCase.input)
		}
	}
}
