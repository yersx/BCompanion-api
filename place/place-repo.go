package place

import "bcompanion/model"

type PlaceRepository interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)

	SavePlace(place model.Place, city string) error
	GetPlaces(city string) ([]*model.Place, error)
	GetPlacesName() ([]*model.Place, error)

	SaveDescription(place model.PlaceDescription) error
	GetDescription(placeName string) (*model.PlaceDescription, error)
}
