# SimpleForecast

SimpleForecast is a command line utility that displays the forecast or the current weather from anywhere in the world inside your terminal.  

## Install

First, you need to add $GOPATH/bin to your path.

```console
$ git clone https://github.com/AlexMercan/SimpleForecast/
$ cd SimpleForecast
$ go install
```
Create an account on [WeatherApi](https://www.weatherapi.com/) to obtain your API key.  
Add that API key as an environment variable with the name MY_WEATHER_API_KEY

## Platforms  

* Windows
* Linux
* macOS

## Usage examples

To display forecast for the next 2 days in New York, in English:

```console
$ simpleforecast -days 2 -location "New York" -lang en
```

#### Output:  
![ForecastUsageSample](https://user-images.githubusercontent.com/35340702/125853410-dd741095-a592-4502-b1fa-7e0550a4678e.png)

To display the current weather at a certain location, in French:

```console
$ simpleforecast -location "Boston" -lang fr
```

#### Output:  
![image](https://user-images.githubusercontent.com/35340702/125855150-f8174e00-a9bb-4ff9-b732-6fb278622916.png)


## Command line options

```
-days     Number of days for the forecast(1 to 3)
-lang     Language for the condition status(sunny, heavy snow, etc.) ( default: en )
-location Location for the weather data(this can be a city name e.g Paris; an ip address e.g:100.0.0.1; latitude and longitude e.g:48.8567,2.3508) (default is location via your ip)
```
## Configuration file

After you run the app for the first time it will create a configuration file called config.json where you can write a default location and language.  
On a **Windows** system this file is located at `%AppData%/simpleforecast/config.json` .  
On a **Unix** system this file is located either at `$XDG_CONFIG_HOME/simpleforecast/config.json` or `$HOME/.config/simpleforecast/config.json` .

