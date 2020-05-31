package hike

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type controller struct{}

var (
	hikeService HikeService
)

type HikeController interface {
	AddHike(w http.ResponseWriter, r *http.Request)
	GetHike(w http.ResponseWriter, r *http.Request)
	GetHikes(w http.ResponseWriter, r *http.Request)
	GetUpcomingHikes(w http.ResponseWriter, r *http.Request)
	GetUpcomingHikesByUser(w http.ResponseWriter, r *http.Request)
	GetUpcomingHikesByPlace(w http.ResponseWriter, r *http.Request)
	GetPastHikesByUser(w http.ResponseWriter, r *http.Request)

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

	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		json.NewEncoder(w).Encode("no token sent")
		w.WriteHeader(404)
		return
	}

	var hike model.Hike
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &hike)
	if err != nil {
		json.NewEncoder(w).Encode("can not get data")
		w.WriteHeader(404)
		return
	}
	s := strings.Split(*hike.StartDate, ".")
	if len(s) != 3 {
		json.NewEncoder(w).Encode("not correct date")
		w.WriteHeader(404)
		return
	}
	startDateStr := s[2] + "-" + s[1] + "-" + s[0]
	layoutISO := "2006-01-02"
	startDate, err := time.Parse(layoutISO, startDateStr)
	if err != nil {
		json.NewEncoder(w).Encode("not correct date")
		w.WriteHeader(404)
		return
	}
	hike.StartDateISO = &startDate

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

	HikeId, ok1 := r.URL.Query()["hike_id"]
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
	if len(hike) < 1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return
}

func (*controller) GetUpcomingHikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	hike, err := hikeService.GetUpcomingHikes()
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	if len(hike) < 1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return
}

func (*controller) GetUpcomingHikesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	hike, err := hikeService.GetUpcomingHikesByUser(token)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	if len(hike) < 1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return
}

func (*controller) GetUpcomingHikesByPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

	Place, ok1 := r.URL.Query()["place_name"]
	if !ok1 || len(Place[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	place := Place[0]

	hike, err := hikeService.GetUpcomingHikesByPlace(place)
	if err != nil || len(hike) < 1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return
}

func (*controller) GetPastHikesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	hike, err := hikeService.GetPastHikesByUser(token)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	if len(hike) < 1 {
		w.WriteHeader(404)
		log.Println("no past hikes")
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(hike)
	return
}

func (*controller) JoinHike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	token := r.Header.Get("Authorization")

	HikeID, ok1 := r.URL.Query()["hike_id"]
	if !ok1 || len(HikeID[0]) < 1 {
		json.NewEncoder(w).Encode("Url Param 'hike_id' is missing")
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
