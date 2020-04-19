package place

import (
	"bcompanion/model"
)

type PlaceService interface {
	SaveCity(city model.City) error
	GetCities() ([]*model.City, error)
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
