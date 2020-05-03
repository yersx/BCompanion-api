package model

type Group struct {
	Name            string           `json:"groupName" bson:"groupName"`
	Description     string           `json:"groupDescription" bson:"groupDescription"`
	Links           string           `json:"groupLinks" bson:"groupLinks"`
	Image           string           `json:"groupPhoto" bson:"groupPhoto"`
	NumberOfMembers int32            `json:"numberOfMembers" bson:"numberOfMembers"`
	NumberOfHikes   int32            `json:"numberOfHikes" bson:"numberOfHikes"`
	CurrentHikes    []*HikeShortInfo `json:"currentHikes" bson:"currentHikes"`
	HikesHistory    []*HikeShortInfo `json:"hikesHistory" bson:"hikesHistory"`
	GroupMedia      []*Media         `json:"groupMedia" bson:"groupMedia"`
	Members         []*Member        `json:"members" bson:"members"`
}

type GroupItem struct {
	Name            string `json:"groupName" bson:"groupName"`
	Description     string `json:"groupDescription" bson:"groupDescription"`
	Image           string `json:"groupPhoto" bson:"groupPhoto"`
	NumberOfMembers int32  `json:"numberOfMembers" bson:"numberOfMembers"`
	NumberOfHikes   int32  `json:"numberOfHikes" bson:"numberOfHikes"`
}

type Member struct {
	Token   string `json:"-" bson:"token"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Photo   string `json:"photo" bson:"photo"`
	Status  string `json:"status" bson:"status"`
	Role    string `json:"role" bson:"role"`
}

type Media struct {
	MediaUrl  string `json:"mediaUrl" bson:"mediaUrl"`
	MediaType string `json:"mediaType" bson:"mediaType"`
}

type HikeShortInfo struct {
	HikePhoto         string `json:"hikePhoto" bson:"hikePhoto"`
	PlaceName         string `json:"placeName" bson:"placeName"`
	StartDate         string `json:"startDate" bson:"startDate"`
	NumberOfMembers   int32  `json:"numberOfMembers" bson:"numberOfMembers"`
	WithOvernightStay bool   `json:"withOvernightStay" bson:"withOvernightStay"`
}
