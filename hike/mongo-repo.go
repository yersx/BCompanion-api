package hike

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

type repo struct{}

func NewMongoRepository() HikeRepository {
	return &repo{}
}

func (*repo) CreateHike(hike model.Hike) string {

	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return "can not find groups collection"
	}

	userCollection, err := db.GetDBCollection("users")
	var user *model.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"phoneNumber", hike.Admins[0]}}).Decode(&user)
	if err != nil {
		return "can not find creater account"
	}

	hike.Members = []*model.Member{
		{
			Token:       user.Token,
			Name:        user.FirstName,
			Surname:     user.LastName,
			Photo:       user.Photo,
			PhoneNumber: user.PhoneNumber,
			Status:      user.Status,
			Role:        "admin",
		},
	}
	hike.HikeID = bson.NewObjectId()

	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"groupName": hike.GroupName},
		bson.D{
			{"$push", bson.D{{"hikesHistory", hike}}},
		},
	)
	log.Printf("added result %v", result)
	if err != nil {
		return "can not create hiking event"
	}
	return ""
}

type fields struct {
	ID              int `bson:"_id"`
	GroupName       int `bson:"groupName"`
	GroupPhoto      int `bson:"groupPhoto"`
	NumberOfMembers int `bson:"numberOfMembers"`
	NumberOfHikes   int `bson:"numberOfHikes"`
}

func (*repo) GetHike(hikeID string) (*model.Hike, error) {

	var hike *model.Hike
	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"members", bson.D{
			{"$elemMatch", bson.D{
				{"hikesHistory", hikeID},
			},
			}},
		},
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&hike)
	if err != nil {
		return nil, err
	}

	hikeNumber := len(hike.Members)

	hike.NumberOfMembers = strconv.Itoa(hikeNumber)

	return hike, nil
}
