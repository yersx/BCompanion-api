package weather

import "bcompanion/model"

type WeatherRepository interface {
	GetPlaceCoordinates(place string) (*model.Coordinate, error)
}
