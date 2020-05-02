package model

type Group struct {
	Name        string `json:"groupName" bson:"groupName"`
	Description string `json:"groupDescription" bson:"groupDescription"`
	Links       string `json:"groupLinks" bson:"groupLinks"`
	Image       string `json:"groupPhoto" bson:"groupPhoto"`
	Owner       string `json:"-" bson:"groupOwner"`
}
