package group

import (
	"bcompanion/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	utils "bcompanion/utils"
)

type controller struct{}

var (
	groupService GroupService
)

type GroupController interface {
	AddGroup(w http.ResponseWriter, r *http.Request)
	GetUserGroups(w http.ResponseWriter, r *http.Request)
	GetAllGroups(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)

	JoinGroup(w http.ResponseWriter, r *http.Request)
	LeaveGroup(w http.ResponseWriter, r *http.Request)
}

// NewPlaceController implements PlaceController
func NewGroupController(service GroupService) GroupController {
	groupService = service
	return &controller{}
}

func (*controller) AddGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group model.Group

	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		json.NewEncoder(w).Encode("no token sent")
		w.WriteHeader(404)
		return
	}

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

	var groupLinks *string
	GroupLinks, ok1 := r.URL.Query()["group_links"]
	if !ok1 || len(GroupLinks[0]) < 1 {

	} else {
		groupLinks = &GroupLinks[0]
	}

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

	// 3. Generate new filename
	nameFile := handler.Filename

	// 4. Read multipart file
	buff, errReadFile := ioutil.ReadAll(file)
	if errReadFile != nil {
		json.NewEncoder(w).Encode("Error reading file")
		w.WriteHeader(404)
		return
	}

	//5. Upload to cloudinary
	resChannelUpload := utils.UploadImage(nameFile, buff)
	cloudinaryInfo := <-resChannelUpload
	close(resChannelUpload)
	if cloudinaryInfo.Err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("internal server error with cloudinary")
		return
	}

	cloudinaryPath := cloudinaryInfo.FilePath

	group = model.Group{
		Name:        groupName,
		Description: &groupDescription,
		Links:       groupLinks,
		Image:       &cloudinaryPath,
	}

	res := groupService.AddGroup(group, token)
	if res != "" {
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode("Successfully added")
	return
}

func (*controller) GetUserGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}

	groups, err := groupService.GetUserGroups(token)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	if len(groups) < 1 {
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(groups)
	return

}

func (*controller) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	groups, err := groupService.GetAllGroups()
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(groups)
	return

}

func (*controller) GetGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var res model.ResponseResult

	GroupName, ok1 := r.URL.Query()["group_name"]
	if !ok1 || len(GroupName[0]) < 1 {
		res.Message = "Url Param 'group_name' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	groupName := GroupName[0]

	group, err := groupService.GetGroup(groupName)
	if err != nil {
		res.Message = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(group)
	return
}

func (*controller) JoinGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	log.Printf("joining group token: %v", token)
	GroupName, ok1 := r.URL.Query()["group_name"]
	if !ok1 || len(GroupName[0]) < 1 {
		json.NewEncoder(w).Encode("Url Param 'group_name' is missing")
		w.WriteHeader(404)
		return
	}
	groupName := GroupName[0]

	response := groupService.JoinGroup(groupName, token)
	if response != "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode("Successfully joined to group")
	return
}

func (*controller) LeaveGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	GroupName, ok1 := r.URL.Query()["group_name"]
	if !ok1 || len(GroupName[0]) < 1 {
		json.NewEncoder(w).Encode("Url Param 'group_name' is missing")
		w.WriteHeader(404)
		w.WriteHeader(404)
		return
	}
	groupName := GroupName[0]

	response := groupService.LeaveGroup(groupName, token)
	if response != "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode("Successfully leaved the group")
	return
}
