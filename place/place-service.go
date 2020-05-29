package place

import (
	"bcompanion/model"
)

type PlaceService interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)
	GetCitiesName() ([]*string, error)
	GetCityCoordinates(city string) ([]*float64, error)

	AddPlace(place model.Place, city string) error
	GetPlaces(city string) ([]*model.Place, error)
	GetPlacesName() ([]*string, error)

	AddPlaceDescription(place model.PlaceDescription) error
	GetPlaceDescription(placeName string) (*model.Description, error)
	GetPlaceRoute(placeName string) (*model.PlaceRoute, error)
}

type service struct{}

var (
	placeRepo PlaceRepository
)

func NewPlaceService(repository PlaceRepository) PlaceService {
	placeRepo = repository
	return &service{}
}

func (*service) SaveCity(city model.City) error {
	return placeRepo.SaveCity(city)
}

func (*service) GetCities() ([]*model.City, error) {
	return placeRepo.GetCities()
}

func (*service) GetCitiesName() ([]*string, error) {
	return placeRepo.GetCitiesName()
}

func (*service) GetCityCoordinates(city string) ([]*float64, error) {
	return placeRepo.GetCityCoordinates(city)
}

func (*service) AddPlace(place model.Place, city string) error {
	return placeRepo.SavePlace(place, city)
}

func (*service) GetPlaces(city string) ([]*model.Place, error) {
	return placeRepo.GetPlaces(city)
}

func (*service) GetPlacesName() ([]*string, error) {
	return placeRepo.GetPlacesName()
}

func (*service) AddPlaceDescription(place model.PlaceDescription) error {
	return placeRepo.SaveDescription(place)
}

func (*service) GetPlaceDescription(placeName string) (*model.Description, error) {
	return placeRepo.GetDescription(placeName)
}

func (*service) GetPlaceRoute(placeName string) (*model.PlaceRoute, error) {
	return placeRepo.GetPlaceRoute(placeName)
}
