package place

import "bcompanion/model"

type PlaceRepository interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)

	SavePlace(city model.Place) error
	GetPlaces() ([]*model.Place, error)
}
