package model

type City struct {
	CityName  string  `json:"cityName" bson:"cityName"`
	CityPhoto string  `json:"cityPhoto" bson:"cityPhoto"`
	Place     []Place `json:"-" bson:"places"`
}

type Place struct {
	PlaceName  string `json:"placeName" bson:"placeName"`
	PlacePhoto string `json:"placePhoto" bson:"placePhoto"`
	CityName   string `json:"cityName" bson:"cityName"`
}

type Places struct {
	Places []*Place `json:"-" bson:"places"`
}

type Cities struct {
	CityNames string `json:"-" bson:"cityName"`
}

type PlaceDescription struct {
	PlaceName        string `json:"placeName" bson:"placeName"`
	PlacePhotos      string `json:"placePhotos" bson:"placePhotos"`
	City             string `json:"city" bson:"city"`
	PlaceDescription string `json:"placeDescription" bson:"placeDescription"`
	Lattitude        string `json:"lat" bson:"lat"`
	Longitude        string `json:"long" bson:"long"`
	// JsonPhoto        string   `json:"placePhotos" bson:"-"`
}

type Description struct {
	PlaceName        string `json:"placeName" bson:"placeName"`
	PlacePhotos      string `json:"placePhotos" bson:"placePhotos"`
	PlaceDescription string `json:"placeDescription" bson:"placeDescription"`
	Lattitude        string `json:"lat" bson:"lat"`
	Longitude        string `json:"long" bson:"long"`
	// JsonPhoto        string   `json:"placePhotos" bson:"-"`
}
