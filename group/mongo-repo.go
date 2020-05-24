package group

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type repo struct{}

func NewMongoRepository() GroupRepository {
	return &repo{}
}

func (*repo) CreateGroup(group model.Group, token string) string {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return "can not find groups collection"
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
		NumberOfMembers: "1",
		NumberOfHikes:   "0",
		Admins:          []string{user.PhoneNumber},
		CurrentHikes:    []*model.Hike{},
		HikesHistory:    []*model.Hike{},
		GroupMedia:      []*model.Media{},
		Members: []*model.Member{
			{
				Token:       token,
				Name:        user.FirstName,
				Surname:     user.LastName,
				Photo:       user.Photo,
				PhoneNumber: user.PhoneNumber,
				Status:      user.Status,
				Role:        "admin",
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

	cursor, err := collection.Find(
		context.TODO(), filter)
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

	b.HikesHistory = GetHike(b.Name)
	numberOfHikes := len(b.HikesHistory) + len(b.CurrentHikes)
	b.NumberOfHikes = strconv.Itoa(numberOfHikes)

	groupNumber := len(b.Members)
	b.NumberOfMembers = strconv.Itoa(groupNumber)

	b.Members = nil
	b.Description = nil
	b.Links = nil
	b.CurrentHikes = nil
	b.Admins = nil
	b.HikesHistory = nil
	b.GroupMedia = nil
	b.Admins = nil
	return b
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

	cursor, err := collection.Find(
		context.TODO(), bson.D{})
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

func (*repo) GetGroup(groupName string) (*model.Group, error) {

	var group *model.Group
	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}

	if groupName != "" {
		err = collection.FindOne(context.TODO(), bson.D{{"groupName", groupName}}).Decode(&group)
		if err != nil {
			return nil, err
		}
	}

	group.HikesHistory = GetHike(groupName)

	numberOfMembers := len(group.Members)
	numberOfHikes := len(group.HikesHistory) + len(group.CurrentHikes)

	if len(group.CurrentHikes) == 0 {
		group.CurrentHikes = nil
	}
	if len(group.HikesHistory) == 0 {
		group.HikesHistory = nil
	}

	group.NumberOfMembers = strconv.Itoa(numberOfMembers)
	group.NumberOfHikes = strconv.Itoa(numberOfHikes)

	return group, nil
}

func (*repo) JoinGroup(groupName string, token string) string {
	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return "can not find groups collection"
	}

	userCollection, err := db.GetDBCollection("users")
	var user *model.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"token", token}}).Decode(&user)
	if err != nil {
		return "can not find  account"
	}
	member := model.Member{
		Token:       token,
		Name:        user.FirstName,
		Surname:     user.LastName,
		Photo:       user.Photo,
		PhoneNumber: user.PhoneNumber,
		Status:      user.Status,
		Role:        "",
	}

	_, err2 := collection.UpdateOne(
		context.TODO(),
		bson.M{"groupName": groupName},
		bson.D{
			{"$push", bson.D{{"members", member}}},
		},
	)
	if err2 != nil {
		return "can not join the group"
	}
	return ""
}

func (*repo) LeaveGroup(groupName string, token string) string {
	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return "can not find groups collection"
	}

	_, err2 := collection.UpdateOne(
		context.TODO(),
		bson.M{"groupName": groupName},
		bson.D{
			{"$pull", bson.D{{"members", bson.M{"token": token}}}},
		},
	)
	if err2 != nil {
		return "can not leave the group"
	}
	return ""
}

func GetHike(groupName string) []*model.Hike {

	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return nil
	}

	cursor, err := collection.Find(
		context.TODO(),
		bson.D{{"groupName", groupName}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.TODO())

	out := make([]*model.Hike, 0)

	for cursor.Next(context.TODO()) {
		hike := new(model.Hike)
		err := cursor.Decode(hike)
		if err != nil {
			return nil
		}
		out = append(out, hike)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}

	return toHikes(out)
}

func toHike(b *model.Hike) *model.Hike {
	numberOfMembers := len(b.Members)
	b.NumberOfMembers = strconv.Itoa(numberOfMembers)
	return b
}

func toHikes(bs []*model.Hike) []*model.Hike {
	out := make([]*model.Hike, len(bs))

	for i, b := range bs {
		out[i] = toHike(b)
	}
	return out
}
