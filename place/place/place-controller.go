package place

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	place "bcompanion/place"
)

type controller struct{}

var (
	placeService place.PlaceService
)

type PlaceController interface {
	AddPlace(w http.ResponseWriter, r *http.Request)
	GetPlaces(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewPlaceController(service place.PlaceService) PlaceController {
	placeService = service
	return &controller{}
}

func (*controller) AddPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var place model.Place
	body, _ := ioutil.ReadAll(r.Body)
	var response model.TokenResult

	err := json.Unmarshal(body, &place)
	if err != nil {
		response.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(400)
		return
	}

	City, ok1 := r.URL.Query()["city"]
	if !ok1 || len(City[0]) < 1 {
		response.Message = "Url Param 'city' is missing"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(404)
		return
	}
	city := City[0]

	err = placeService.AddPlace(place, city)
	if err != nil {
		response.Message = "Can not add place"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(404)
		return
	}

	response.Message = "Successfully added"
	json.NewEncoder(w).Encode(response)
	return
}

func (*controller) GetPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res model.ResponseResult

	City, ok1 := r.URL.Query()["city_name"]
	if !ok1 || len(City[0]) < 1 {
		res.Message = "Url Param 'city_name' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	city := City[0]

	cities, err := placeService.GetPlaces(city)
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}
