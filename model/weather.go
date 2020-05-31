package model

type Coordinate struct {
	Lattitude *float64 `json:"latitude" bson:"latitude"`
	Longitude *float64 `json:"longitude" bson:"longitude"`
}
