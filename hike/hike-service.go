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
	GetUpcomingHikesByPlace(place string) ([]*model.Hike, error)

	JoinHike(hikeId string, token string) string
	LeaveHike(hikeId string, token string) string
}

type service struct{}

var (
	hikeRepo HikeRepository
)

var hikeImages = map[string]string{
	"Шарын":    "https://i.ibb.co/s1DnwVM/sharyn.jpg",
	"Колсай":   "https://i.ibb.co/h1srvX0/kolsay.jpg",
	"Капшагай": "https://i.ibb.co/X3MHdQF/kapshagay.png",
	"Высокогорный курорт Ак-Булак":      "https://i.imgur.com/SAXZWq6.jpg",
	"Озеро М. Маметовой":                "https://i.imgur.com/A8ocIYR.jpg",
	"Плато Устюрт":                      "https://i.imgur.com/IAYp3Qi.jpg",
	"скала Шеркала":                     "https://i.imgur.com/0EhqJiF.jpg",
	"Большое Алматинское Озеро":         "https://i.ibb.co/VQBrxsT/bao.jpg",
	"Бутаковский водопад":               "https://i.ibb.co/RvZk2yH/butakovka.jpg",
	"Кок-жайлау":                        "https://i.imgur.com/dtjCydx.jpg",
	"Качели":                            "https://x-travels.ru/wp-content/uploads/2019/11/IMG_6168.jpg",
	"Мынжылкы":                          "https://i.ibb.co/t490TxK/mynzhylky.jpg",
	"Водопад Горельник":                 "https://i.imgur.com/HMSt2Kk.jpg",
	"Водопад Чукур":                     "https://i.ibb.co/6NxqFxY/chukur.jpg",
	"Пик Фурманова":                     "https://i.ibb.co/RpmYCv4/furmanova.jpg",
	"Озеро Титова":                      "https://i.ibb.co/YN6NWQz/titova.jpg",
	"Пещера Акмешит":                    "https://i.imgur.com/c2hqwHn.jpg",
	"Сайрам-Огемский национальный парк": "https://i.imgur.com/s04vaJL.jpg",
	"Казыгурт":                          "https://i.imgur.com/E8JK8xd.jpg",
	"Городище Сауран":                   "https://i.imgur.com/mYrGfGm.jpg",
}

func NewHikeService(repository HikeRepository) HikeService {
	hikeRepo = repository
	return &service{}
}

func (*service) AddHike(hike model.Hike, token string) string {

	photo := hikeImages[*hike.PlaceName]
	hike.HikePhoto = &photo
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

func (*service) GetUpcomingHikesByPlace(place string) ([]*model.Hike, error) {
	return hikeRepo.GetUpcomingHikesByPlace(place)
}

func (*service) GetUpcomingHikesByUser(token string) ([]*model.Hike, error) {
	return hikeRepo.GetUpcomingHikesByUser(token)
}

func (*service) GetPastHikesByUser(token string) ([]*model.Hike, error) {
	return hikeRepo.GetPastHikesByUser(token)
}
