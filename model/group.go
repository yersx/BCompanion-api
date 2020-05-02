package model

type Group struct {
	Name            string `json:"groupName" bson:"groupName"`
	Description     string `json:"-" bson:"groupDescription"`
	Links           string `json:"-" bson:"groupLinks"`
	Image           string `json:"groupPhoto" bson:"groupPhoto"`
	Owner           string `json:"-" bson:"groupOwner"`
	NumberOfMembers string `json:"numberOfMembers" bson:"numberOfMembers"`
	NumberOfHikes   string `json:"numberOfHikes" bson:"numberOfHikes"`
}
