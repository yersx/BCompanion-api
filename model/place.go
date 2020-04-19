package model

type City struct {
	CityName  string `json:"cityName" bson:"cityName"`
	CityPhoto string `json:"cityPhoto" bson:"cityPhoto"`
}
