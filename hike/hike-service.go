package hike

import (
	"bcompanion/model"
)

type HikeService interface {
	AddHike(hike model.Hike, token string) string
	GetHike(hikeId string) (*model.Hike, error)
}

type service struct{}

var (
	hikeRepo HikeRepository
)

func NewHikeService(repository HikeRepository) HikeService {
	hikeRepo = repository
	return &service{}
}

func (*service) AddHike(hike model.Hike, token string) string {
	return hikeRepo.CreateHike(hike, token)
}

func (*service) GetHike(hikeId string) (*model.Hike, error) {
	return hikeRepo.GetHike(hikeId)
}
