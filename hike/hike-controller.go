package hike

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type controller struct{}

var (
	hikeService HikeService
)

type HikeController interface {
	AddHike(w http.ResponseWriter, r *http.Request)
	GetHike(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewHikeController(service HikeService) HikeController {
	hikeService = service
	return &controller{}
}

func (*controller) AddHike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var hike model.Hike
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &hike)
	if err != nil {
		json.NewEncoder(w).Encode("can not get data")
		w.WriteHeader(400)
		return
	}

	res := hikeService.AddHike(hike)
	if res != "" {
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode("Successfully added")
	return
}

func (*controller) GetHike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	HikeId, ok1 := r.URL.Query()["hikeId"]
	if !ok1 || len(HikeId[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	hikeId := HikeId[0]

	hike, err := hikeService.GetHike(hikeId)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return

}
