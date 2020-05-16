package group

import "bcompanion/model"

type HikeRepository interface {
	CreateHike(hike model.Hike) string
	GetHike(hikeId string) (*model.Hike, error)
}
