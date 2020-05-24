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
	GetPlacesName(w http.ResponseWriter, r *http.Request)
	AddPlaceDescription(w http.ResponseWriter, r *http.Request)
	GetPlaceDescription(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewPlaceController(service place.PlaceService) PlaceController {
	placeService = service
	return &controller{}
}

func (*controller) AddPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var res model.ResponseResult

	City, ok1 := r.URL.Query()["city_name"]
	if !ok1 || len(City[0]) < 1 {
		res.Message = "Url Param 'city_name' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	city := City[0]

	places, err := placeService.GetPlaces(city)
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(places)
	return

}

func (*controller) GetPlacesName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var res model.ResponseResult

	cities, err := placeService.GetPlacesName()
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}

func (*controller) AddPlaceDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var place model.PlaceDescription
	body, _ := ioutil.ReadAll(r.Body)
	var response model.TokenResult

	err := json.Unmarshal(body, &place)
	if err != nil {
		response.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(400)
		return
	}

	err = placeService.AddPlaceDescription(place)
	if err != nil {
		response.Message = "Can not add description"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(404)
		return
	}

	response.Message = "Successfully added"
	json.NewEncoder(w).Encode(response)
	return
}

func (*controller) GetPlaceDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var res model.ResponseResult

	Place, ok1 := r.URL.Query()["place_name"]
	if !ok1 || len(Place[0]) < 1 {
		res.Message = "Url Param 'place_name' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	place := Place[0]

	placeDescription, err := placeService.GetPlaceDescription(place)
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	description := model.Description{
		PlaceName:        placeDescription.PlaceName,
		PlacePhotos:      placeDescription.PlacePhotos,
		PlaceDescription: placeDescription.PlaceDescription,
		Lattitude:        placeDescription.Lattitude,
		Longitude:        placeDescription.Longitude,
	}

	json.NewEncoder(w).Encode(description)
	return

}
