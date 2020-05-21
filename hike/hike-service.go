package hike

import (
	"bcompanion/model"
)

type HikeService interface {
	AddHike(hike model.Hike, token string) string
	GetHike(hikeId string) (*model.Hike, error)

	JoinHike(hikeId string, token string) string
	LeaveHike(hikeId string, token string) string
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

func (*service) JoinHike(hikeId string, token string) string {
	return hikeRepo.JoinHike(hikeId, token)
}

func (*service) LeaveHike(hikeId string, token string) string {
	return hikeRepo.LeaveHike(hikeId, token)
}
