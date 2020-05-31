package model

type Coordinate struct {
	Lattitude *float64 `json:"latitude" bson:"latitude"`
	Longitude *float64 `json:"longitude" bson:"longitude"`
}

type Weather struct {
	Daily []struct {
		Date      int64   `json:"dt"`
		Humidity  int     `json:"humidity"`
		WindSpeed float64 `json:"wind_speed"`
		Clouds    int     `json:"clouds"`
		Temp      struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
		} `json:"temp"`
		Weather []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"daily"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

type WeatherDay struct {
	PlaceName   string `json:"placeName"`
	Date        string `json:"date"`
	Day         string `json:"day"`
	Image       string `json:"image"`
	DayDegree   string `json:"dayDegree"`
	NightDegree string `json:"nightDegree"`
	Humidity    string `json:"humidity"`
	WindSpeed   string `json:"wind_speed"`
	Clouds      string `json:"clouds"`
}
