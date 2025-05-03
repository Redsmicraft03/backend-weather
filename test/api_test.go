package test

import (
	"testing"

	"github.com/Redsmicraft03/backend-weather/api"
	"github.com/Redsmicraft03/backend-weather/entity"
)

func TestGetWeather(t *testing.T) {
	api.GetWeather(entity.City{
			NameCity string `json:"name_city"`
		}{
			{
				NameCity: "Bogor",
			},
			{
				NameCity: "Bandung",
			},
			{
				NameCity: "Gorontalo",
			},
		},Location: []struct {
		
	})
}
