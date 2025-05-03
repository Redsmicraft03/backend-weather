package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Redsmicraft03/backend-weather/api"
	"github.com/Redsmicraft03/backend-weather/entity"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file.json>")
		return
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var city entity.City
	err = json.Unmarshal(data, &city)
	if err != nil {
		panic(err)
	}

	api.GetWeather(city)
}
