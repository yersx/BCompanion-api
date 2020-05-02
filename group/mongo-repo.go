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

func (*repo) CreateGroup(group model.Group, token string) error {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return err
	}

	group = model.Group{
		Name:        group.Name,
		Description: group.Description,
		Links:       group.Links,
		Image:       group.Image,
		Owner:       token,
	}

	_, err = collection.InsertOne(context.TODO(), group)
	// Check if Group Insertion Fails
	if err != nil {
		return err
	}
	return nil
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
