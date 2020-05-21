package hike

import "bcompanion/model"

type HikeRepository interface {
	CreateHike(hike model.Hike, token string) string
	GetHike(hikeId string) (*model.Hike, error)

	JoinHike(hikeId string, token string) string
	LeaveHike(hikeId string, token string) string
}
