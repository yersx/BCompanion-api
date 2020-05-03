package group

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type fields struct {
	ID              int `bson:"_id"`
	GroupName       int `bson:"groupName"`
	GroupPhoto      int `bson:"groupPhoto"`
	NumberOfMembers int `bson:"numberOfMembers"`
	NumberOfHikes   int `bson:"numberOfHikes"`
}

func (*repo) GetUserGroups(token string) ([]*model.Group, error) {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"members", bson.D{
			{"$elemMatch", bson.D{
				{"token", token},
			},
			}},
		},
	}

	projection := fields{
		ID:              0,
		GroupName:       1,
		GroupPhoto:      1,
		NumberOfMembers: 1,
		NumberOfHikes:   1,
	}

	cursor, err := collection.Find(
		context.TODO(), filter,
		options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	out := make([]*model.Group, 0)

	for cursor.Next(context.TODO()) {
		group := new(model.Group)
		err := cursor.Decode(group)
		if err != nil {
			return nil, err
		}
		out = append(out, group)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toGroups(out), nil
}

func toGroup(b *model.Group) *model.Group {
	return &model.Group{
		Name:            b.Name,
		Description:     b.Description,
		Image:           b.Image,
		NumberOfHikes:   b.NumberOfHikes,
		NumberOfMembers: b.NumberOfMembers,
	}
}

func toGroups(bs []*model.Group) []*model.Group {
	out := make([]*model.Group, len(bs))

	for i, b := range bs {
		out[i] = toGroup(b)
	}
	return out
}

func (*repo) GetAllGroups() ([]*model.Group, error) {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})
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
