package hike

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type controller struct{}

var (
	hikeService HikeService
)

type HikeController interface {
	AddHike(w http.ResponseWriter, r *http.Request)
	GetHike(w http.ResponseWriter, r *http.Request)

	JoinHike(w http.ResponseWriter, r *http.Request)
	LeaveHike(w http.ResponseWriter, r *http.Request)
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

	token := r.Header.Get("Authorization")

	var hike model.Hike
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &hike)
	if err != nil {
		json.NewEncoder(w).Encode("can not get data")
		w.WriteHeader(404)
		return
	}

	res := hikeService.AddHike(hike, token)
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

func (*controller) GetHikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var groupName = ""
	GroupName, ok1 := r.URL.Query()["group_name"]
	if !ok1 || len(GroupName[0]) < 1 {

	} else {
		groupName = GroupName[0]
	}

	hike, err := hikeService.GetHikes(groupName)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return

}

func toHike(b *model.Hike) *model.Hike {
	numberOfMembers := len(b.Members)
	b.NumberOfMembers = strconv.Itoa(numberOfMembers)
	return b
}

func toHikes(bs []*model.Hike) []*model.Hike {
	out := make([]*model.Hike, len(bs))

	for i, b := range bs {
		out[i] = toHike(b)
	}
	return out
}

func (*controller) JoinHike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token := r.Header.Get("Authorization")

	HikeID, ok1 := r.URL.Query()["hike_id"]
	if !ok1 || len(HikeID[0]) < 1 {
		json.NewEncoder(w).Encode("Url Param 'hike_id' is missing")
		w.WriteHeader(404)
		return
	}
	hikeID := HikeID[0]

	response := hikeService.JoinHike(hikeID, token)
	if response != "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode("Successfully joined to hiking event")
	return
}

func (*controller) LeaveHike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token := r.Header.Get("Authorization")

	HikeID, ok1 := r.URL.Query()["hike_id"]
	if !ok1 || len(HikeID[0]) < 1 {
		json.NewEncoder(w).Encode("Url Param 'hike_id' is missing")
		w.WriteHeader(404)
		w.WriteHeader(404)
		return
	}
	hikeID := HikeID[0]

	response := hikeService.LeaveHike(hikeID, token)
	if response != "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode("Successfully leaved the hiking event")
	return
}
