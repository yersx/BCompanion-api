package weather

import (
	"encoding/json"
	"net/http"
)

type controller struct{}

var (
	weatherService WeatherService
)

type WeatherController interface {
	GetWeekWeather(w http.ResponseWriter, r *http.Request)
	GetDayWeather(w http.ResponseWriter, r *http.Request)
}

func NewWeatherController(service WeatherService) WeatherController {
	weatherService = service
	return &controller{}
}

func (*controller) GetWeekWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	City, ok1 := r.URL.Query()["place_name"]
	if !ok1 || len(City[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	city := City[0]

	weather, err := weatherService.GetWeekWeather(city)
	if err != nil || weather == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(weather)
	return

}

func (*controller) GetDayWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	PlaceName, ok1 := r.URL.Query()["place_name"]
	if !ok1 || len(PlaceName[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	place := PlaceName[0]

	Date, ok1 := r.URL.Query()["date"]
	if !ok1 || len(Date[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	date := Date[0]

	weather, err := weatherService.GetDayWeather(place, date)
	if err != nil || weather == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(weather)
	return

}
