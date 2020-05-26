package hike

import (
	"bcompanion/model"
)

type HikeService interface {
	AddHike(hike model.Hike, token string) string
	GetHike(hikeId string) (*model.Hike, error)

	GetHikes(groupName string) ([]*model.Hike, error)
	GetUpcomingHikes() ([]*model.Hike, error)
	GetUpcomingHikesByUser(token string) ([]*model.Hike, error)
	GetPastHikesByUser(token string) ([]*model.Hike, error)

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

func (*service) GetHikes(groupName string) ([]*model.Hike, error) {
	return hikeRepo.GetHikes(groupName)
}

func (*service) GetUpcomingHikes() ([]*model.Hike, error) {
	return hikeRepo.GetUpcomingHikes()
}

func (*service) GetUpcomingHikesByUser(token string) ([]*model.Hike, error) {
	return hikeRepo.GetUpcomingHikesByUser(token)
}

func (*service) GetPastHikesByUser(token string) ([]*model.Hike, error) {
	return hikeRepo.GetPastHikesByUser(token)
}
