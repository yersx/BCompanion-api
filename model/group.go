package model

type Group struct {
	Name        string `json:"group_name" bson:"group_name"`
	Description string `json:"group_description" bson:"group_description"`
	Links       string `json:"group_links" bson:"group_links"`
	Image       string `json:"-" bson:"group_photo"`
	Owner       string `json:"-" bson:"group_owner"`
}
