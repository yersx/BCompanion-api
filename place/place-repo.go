package place

import "bcompanion/model"

type PlaceRepository interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)
	GetCitiesName() ([]*string, error)
	GetCityCoordinates(city string) ([]*float64, error)

	SavePlace(place model.Place, city string) error
	GetPlaces(city string) ([]*model.Place, error)
	GetPlacesName() ([]*string, error)

	SaveDescription(place model.PlaceDescription) error
	GetDescription(placeName string) (*model.Description, error)
}
