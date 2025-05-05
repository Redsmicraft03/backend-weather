package entity

type City struct {
	Location []struct {
		NameCity string `json:"name_city"`
	} `json:"Location"`
}

type WeatherData struct {
	Name string `json:"name"`
	Error string `json:"error,omitempty"`
	Main struct {
		Temp float64 `json:"temp,omitempty"`
		Humidity int `json:"humidity,omitempty"` 
	} `json:"main,omitempty"`
	Weather []struct {
		Description string `json:"description,omitempty"`
	} `json:"weather,omitempty"`
	TimeZone int `json:"timezone,omitempty"`
	GMT string `json:"gmt,omitempty"`
}
