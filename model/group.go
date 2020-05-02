package model

type Group struct {
	Name            string    `json:"groupName" bson:"groupName"`
	Description     string    `json:"-" bson:"groupDescription"`
	Links           string    `json:"-" bson:"groupLinks"`
	Image           string    `json:"groupPhoto" bson:"groupPhoto"`
	Owner           string    `json:"-" bson:"groupOwner"`
	NumberOfMembers string    `json:"numberOfMembers" bson:"numberOfMembers"`
	NumberOfHikes   string    `json:"numberOfHikes" bson:"numberOfHikes"`
	Members         []*Member `json:"-" bson:"members"`
}

type Member struct {
	Token   string `json:"-" bson:"token"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Photo   string `json:"photo" bson:"photo"`
	Status  string `json:"status" bson:"status"`
	Role    string `json:"role" bson:"role"`
}
