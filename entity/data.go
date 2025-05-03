package entity

type City struct {
	Location []struct {
		NameCity string `json:"name_city"`
	} `json:"Location"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
		Humidity int `json:"humidity"` 
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	TimeZone int `json:"timezone"`
	GMT string
}
