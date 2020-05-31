package model

type Coordinate struct {
	Lattitude *float64 `json:"latitude" bson:"latitude"`
	Longitude *float64 `json:"longitude" bson:"longitude"`
}

type Weather struct {
	Daily []struct {
		Date int64 `json:"dt"`
		Temp struct {
			Morning float64 `json:"morn"`
			Night   float64 `json:"night"`
		} `json:"temp"`
		Weather []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"daily"`
	Timezone string `json:"timezone"`
}

type WeatherDay struct {
	PlaceName   string `json:"placeName"`
	Date        string `json:"date"`
	Day         string `json:"day"`
	Image       string `json:"image"`
	DayDegree   string `json:"dayDegree"`
	NightDegree string `json:"nightDegree"`
}
