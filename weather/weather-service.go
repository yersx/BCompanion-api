package weather

import "bcompanion/model"

type WeatherService interface {
	GetWeekWeather(place string) (*model.Coordinate, error)
}

type service struct{}

var (
	weatherRepo WeatherRepository
)

func NewWeatherService(repository WeatherRepository) WeatherService {
	weatherRepo = repository
	return &service{}
}

func (*service) GetWeekWeather(place string) (*model.Coordinate, error) {
	return weatherRepo.GetPlaceCoordinates(place)
}
