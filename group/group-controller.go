package group

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"log"
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

	GroupName, ok1 := r.URL.Query()["group_name"]
	if !ok1 || len(GroupName[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	groupName := GroupName[0]

	GroupDescription, ok1 := r.URL.Query()["group_description"]
	if !ok1 || len(GroupDescription[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	groupDescription := GroupDescription[0]

	GroupLinks, ok1 := r.URL.Query()["group_links"]
	if !ok1 || len(GroupLinks[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	groupLinks := GroupLinks[0]

	err := r.ParseMultipartForm(0)
	if err != nil {
		json.NewEncoder(w).Encode("Does not have image")
		w.WriteHeader(404)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		json.NewEncoder(w).Encode("Error retrieving image")
		w.WriteHeader(404)
		return
	}
	defer file.Close()
	log.Printf("fileName %+v\n", handler.Filename)
	log.Printf("fileName %+v\n", handler.Size)
	log.Printf("fileName %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		log.Printf("Error creating temp image")
		json.NewEncoder(w).Encode("Error creating temp image")
		w.WriteHeader(404)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		json.NewEncoder(w).Encode("Error reading file")
		w.WriteHeader(404)
		return
	}
	tempFile.Write(fileBytes)

	log.Output(1, "successfully uploaded")

	token := r.Header.Get("Authorization")

	group = model.Group{
		Name:        groupName,
		Description: groupDescription,
		Links:       groupLinks,
		Image:       tempFile.Name(),
		Owner:       token,
	}

	err = groupService.AddGroup(group)
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
