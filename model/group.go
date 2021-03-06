package model

type Group struct {
	Name            string    `json:"groupName" bson:"groupName"`
	Description     *string   `json:"groupDescription" bson:"groupDescription"`
	Links           *string   `json:"groupLinks" bson:"groupLinks"`
	Image           *string   `json:"groupPhoto" bson:"groupPhoto"`
	NumberOfMembers string    `json:"numberOfMembers" bson:"numberOfMembers"`
	NumberOfHikes   string    `json:"numberOfHikes" bson:"numberOfHikes"`
	CurrentHikes    []*Hike   `json:"upcomingHikes" bson:"upcomingHikes"`
	Admins          []string  `json:"admins" bson:"admins"`
	HikesHistory    []*Hike   `json:"hikesHistory" bson:"hikesHistory"`
	GroupMedia      []*Media  `json:"groupMedia,omitempty" bson:"groupMedia"`
	Members         []*Member `json:"members" bson:"members"`
}

type Member struct {
	Token       string `json:"-" bson:"token"`
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Photo       string `json:"photo" bson:"photo"`
	Status      string `json:"status" bson:"status"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	Role        string `json:"role" bson:"role"`
}

type Media struct {
	MediaUrl  string `json:"mediaUrl" bson:"mediaUrl"`
	MediaType string `json:"mediaType" bson:"mediaType"`
}
