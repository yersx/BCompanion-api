package place

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type controller struct{}

var (
	placeService PlaceService
)

type CityController interface {
	AddCity(w http.ResponseWriter, r *http.Request)
	GetCities(w http.ResponseWriter, r *http.Request)
	GetCitiesName(w http.ResponseWriter, r *http.Request)
	GetCityCoordinates(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewCityController(service PlaceService) CityController {
	placeService = service
	return &controller{}
}

func (*controller) AddCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var city model.City
	body, _ := ioutil.ReadAll(r.Body)
	var response model.TokenResult

	err := json.Unmarshal(body, &city)
	if err != nil {
		response.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(400)
		return
	}

	err = placeService.SaveCity(city)
	if err != nil {
		response.Message = "Can not add city"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(404)
		return
	}

	response.Message = "Successfully added"
	json.NewEncoder(w).Encode(response)
	return
}

func (*controller) GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var res model.ResponseResult

	cities, err := placeService.GetCities()
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}

func (*controller) GetCitiesName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	cities, err := placeService.GetCitiesName()
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}

func (*controller) GetCityCoordinates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	City, ok1 := r.URL.Query()["city_name"]
	if !ok1 || len(City[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	city := City[0]

	cities, err := placeService.GetCityCoordinates(city)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}
