package group

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type controller struct{}

var (
	groupService GroupService
)

type GroupController interface {
	AddGroup(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewGroupController(service GroupService) GroupController {
	groupService = service
	return &controller{}
}

func (*controller) AddGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group model.Group
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &group)
	if err != nil {
		json.NewEncoder(w).Encode("No Fields Were Sent In")
		w.WriteHeader(404)
		return
	}

	token := r.Header.Get("Authorization")

	err = groupService.AddGroup(group, token)
	if err != nil {
		json.NewEncoder(w).Encode("Can not add group")
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode("Successfully added")
	return
}

func (*controller) GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	cities, err := groupService.GetGroups(token)
	if err != nil {
		json.NewEncoder(w).Encode("can not get groups")
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(cities)
	return

}
