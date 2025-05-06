package test

import (
	"testing"

	"github.com/Redsmicraft03/backend-weather/api"
	"github.com/Redsmicraft03/backend-weather/entity"
)

func TestGetWeather(t *testing.T) {
	api.GetWeather(entity.City{
		Location: []struct {
			NameCity string `json:"name_city"`
		}{
			{
				NameCity: "Purwakarta",
			},
			{
				NameCity: "Purbalingga",
			},
			{
				NameCity: "",
			},
			{
				NameCity: "qwert",
			},
		},
	})
}


func TestBlankCity(t *testing.T)  {
	api.GetWeather(entity.City{
		Location: []struct {
			NameCity string `json:"name_city"`
		}{
			{
				NameCity: "",
			},
			{
				NameCity: "",
			},
			{
				NameCity: "",
			},
		},
	})
}

func TestNotFoundCity(t *testing.T)  {
	api.GetWeather(entity.City{
		Location: []struct {
			NameCity string `json:"name_city"`
		}{
			{
				NameCity: "qwerty",
			},
			{
				NameCity: "gotham",
			},
			{
				NameCity: "HanifEly",
			},
		},
	})
}