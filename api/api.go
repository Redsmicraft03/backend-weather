package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Redsmicraft03/backend-weather/entity"
	"github.com/joho/godotenv"
)

func GetWeather(city entity.City) {
	log.SetFlags(log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("WEATHER_API_KEY")

	var urls []string
	var weatherData entity.WeatherData
	var allWeatherData []entity.WeatherData
	var gmt int

	for _, city := range city.Location {
		city.NameCity = url.QueryEscape(city.NameCity)
		urls = append(urls, fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city.NameCity, apiKey))
	}

	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&weatherData)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == http.StatusNotFound {
			weatherData := entity.WeatherData{
				Name:  city.Location[i].NameCity,
				Error: "City not found",
			}
			allWeatherData = append(allWeatherData, weatherData)
		} else if city.Location[i].NameCity == "" {
			weatherData := entity.WeatherData{
				Name:  "isi kota nya bambang",
				Error: "City not found",
			}
			allWeatherData = append(allWeatherData, weatherData)
		} else {
			gmt = weatherData.TimeZone / 3600
			if gmt < 0 {
				weatherData.GMT = fmt.Sprintf("GMT%d", gmt)
			} else {
				weatherData.GMT = fmt.Sprintf("GMT+%d", gmt)
			}
			allWeatherData = append(allWeatherData, weatherData)
		}

	}

	file, err := os.Create("response.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(allWeatherData)
	if err != nil {
		log.Fatal(err)
	}
}
