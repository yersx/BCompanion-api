package group

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type repo struct{}

func NewMongoRepository() GroupRepository {
	return &repo{}
}

func (*repo) CreateGroup(group model.Group, token string) string {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return "can not fing groups collection"
	}

	userCollection, err := db.GetDBCollection("users")
	var user *model.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"token", token}}).Decode(&user)
	if err != nil {
		return "can not find creater account"
	}

	group = model.Group{
		Name:            group.Name,
		Description:     group.Description,
		Links:           group.Links,
		Image:           group.Image,
		NumberOfMembers: 1,
		NumberOfHikes:   0,
		CurrentHikes:    []*model.HikeShortInfo{},
		HikesHistory:    []*model.HikeShortInfo{},
		GroupMedia:      []*model.Media{},
		Members: []*model.Member{
			{
				Token:   token,
				Name:    user.FirstName,
				Surname: user.LastName,
				Photo:   user.Photo,
				Status:  user.Status,
				Role:    "admin",
			},
		},
	}

	var result model.Group
	err = collection.FindOne(context.TODO(), bson.D{{"groupName", group.Name}}).Decode(&result)
	if err != nil {

		_, err = collection.InsertOne(context.TODO(), group)
		// Check if Group Insertion Fails
		if err != nil {
			return "can not add"
		}
		return ""
	} else {
		return "already existed"
	}
}

func (*repo) GetGroups(token string) ([]*model.Group, error) {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{{"group_owner", token}})
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	out := make([]*model.City, 0)

	for cursor.Next(context.TODO()) {
		city := new(model.City)
		err := cursor.Decode(city)
		if err != nil {
			return nil, err
		}

		out = append(out, city)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
