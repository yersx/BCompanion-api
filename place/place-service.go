package place

import (
	"bcompanion/model"
)

type PlaceService interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)

	AddPlace(place model.Place, city string) error
	GetPlaces(city string) ([]*model.Place, error)
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

func (*service) AddPlace(place model.Place, city string) error {
	return placeRepo.SavePlace(place, city)
}

func (*service) GetPlaces(city string) ([]*model.Place, error) {
	return placeRepo.GetPlaces(city)
}