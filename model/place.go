package model

type City struct {
	CityName  string  `json:"cityName" bson:"cityName"`
	CityPhoto string  `json:"cityPhoto" bson:"cityPhoto"`
	Place     []Place `json:"-" bson:"place"`
}

type Place struct {
	PlaceName  string `json:"placeName" bson:"placeName"`
	PlacePhoto string `json:"placePhoto" bson:"placePhoto"`
	CityName   string `json:"cityName" bson:"cityName"`
}
