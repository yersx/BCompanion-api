package hike

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"
	"strconv"

	bsonmongo "go.mongodb.org/mongo-driver/bson"
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
	err = userCollection.FindOne(context.TODO(), bsonmongo.D{{"phoneNumber", hike.Admins[0]}}).Decode(&user)
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
	log.Println("NewObjectId: %s" + hike.HikeID)

	_, err2 := collection.UpdateOne(
		context.TODO(),
		bsonmongo.M{"groupName": hike.GroupName},
		bsonmongo.D{
			{"$push", bsonmongo.D{{"hikesHistory", hike}}},
		},
	)
	if err2 != nil {
		return "can not create hiking event"
	}
	return ""
}

func (*repo) GetHike(hikeID string) (*model.Hike, error) {

	var hike *model.Hike
	collection, err := db.GetDBCollection("groups")
	if err != nil {
		return nil, err
	}
	filter := bsonmongo.D{{"groupName", "Campit"}}

	err = collection.FindOne(context.TODO(), filter).Decode(&hike)
	if err != nil {
		return nil, err
	}

	hikeNumber := len(hike.Members)

	hike.NumberOfMembers = strconv.Itoa(hikeNumber)

	return hike, nil
}
