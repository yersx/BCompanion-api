package hike

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"
	"strconv"
	"time"

	bsonmongo "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct{}

func NewMongoRepository() HikeRepository {
	return &repo{}
}

func (*repo) CreateHike(hike model.Hike, token string) string {

	userCollection, err := db.GetDBCollection("users")
	var user *model.User
	err = userCollection.FindOne(context.TODO(), bsonmongo.D{{"token", token}}).Decode(&user)
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

	layoutISO := "2020-12-20"
	startDate, err := time.Parse(layoutISO, *hike.StartDate)
	if err != nil {
		return "not correct date"
	}
	hike.StartDateISO = &startDate

	hike.Admins = []*string{
		&user.PhoneNumber,
	}

	hike.HikeID = primitive.NewObjectID()
	hikesCollection, _ := db.GetDBCollection("hikes")
	_, err = hikesCollection.InsertOne(context.TODO(), hike)
	if err != nil {
		return "can not add hike"
	}

	// collection, err := db.GetDBCollection("groups")
	// if err != nil {
	// 	return "can not find groups collection"
	// }
	// _, err2 := collection.UpdateOne(
	// 	context.TODO(),
	// 	bsonmongo.M{"groupName": hike.GroupName},
	// 	bsonmongo.D{
	// 		{"$push", bsonmongo.D{{"hikesHistoryRef", hike.HikeID}}},
	// 	},
	// )
	// if err2 != nil {
	// 	return "can not create hiking event"
	// }
	return ""
}

func (*repo) GetHike(hikeID string) (*model.Hike, error) {

	var hike *model.Hike
	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return nil, err
	}

	id, _ := primitive.ObjectIDFromHex(hikeID)

	log.Println("hikeId:" + hikeID)
	err = collection.FindOne(context.TODO(), bsonmongo.M{"_id": id}).Decode(&hike)
	if err != nil {
		return nil, err
	}

	log.Println("user: %s", hike)

	hikeNumber := len(hike.Members)

	hike.NumberOfMembers = strconv.Itoa(hikeNumber)

	return hike, nil
}

func (*repo) GetHikes(groupName string) ([]*model.Hike, error) {

	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return nil, err
	}
	filter := bsonmongo.D{{}}
	if groupName != "" {
		filter = bsonmongo.D{{"groupName", groupName}}
	}

	cursor, err := collection.Find(
		context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	out := make([]*model.Hike, 0)

	for cursor.Next(context.TODO()) {
		hike := new(model.Hike)
		err := cursor.Decode(hike)
		if err != nil {
			return nil, err
		}
		out = append(out, hike)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toHikes(out), nil
}

var projection = bsonmongo.D{
	{"hikeId", 1},
	{"groupName", 1},
	{"groupPhoto", 1},
	{"placeName", 1},
	{"hikePhoto", 1},
	{"hikeByCar", 1},
	{"withOvernightStay", 1},
	{"startDate", 1},
	{"gatheringCity", 1},
	{"numberOfMembers", 1},
	{"members", 1},
}

func (*repo) GetUpcomingHikes() ([]*model.Hike, error) {

	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(
		context.TODO(),
		bsonmongo.D{},
		options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	out := make([]*model.Hike, 0)

	for cursor.Next(context.TODO()) {
		hike := new(model.Hike)
		err := cursor.Decode(hike)
		if err != nil {
			return nil, err
		}
		out = append(out, hike)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toHikes(out), nil
}

func toHike(b *model.Hike) *model.Hike {
	numberOfMembers := len(b.Members)
	b.NumberOfMembers = strconv.Itoa(numberOfMembers)
	b.Members = nil
	return b
}

func toHikes(bs []*model.Hike) []*model.Hike {
	out := make([]*model.Hike, len(bs))

	for i, b := range bs {
		out[i] = toHike(b)
	}
	return out
}

// func GetHike(hikeID string) []*model.Hike {
// 	hikesCollection, _ := db.GetDBCollection("hikes")
// 	cursor, err := hikesCollection.Find(
// 		context.TODO(),
// 		bson.M{"_id": hikeID})
// 	if err != nil {
// 		return nil
// 	}
// 	defer cursor.Close(context.TODO())

// 	out := make([]*model.Hike, 0)

// 	for cursor.Next(context.TODO()) {
// 		hike := new(model.Hike)
// 		err := cursor.Decode(hike)
// 		if err != nil {
// 			return nil
// 		}
// 		out = append(out, hike)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil
// 	}

// 	return toHikes(out)
// }

// func toHike(b *model.Hike) *model.Hike {
// 	numberOfMembers := len(b.Members)
// 	b.NumberOfMembers = strconv.Itoa(numberOfMembers)
// 	return b
// }

// func toHikes(bs []*model.Hike) []*model.Hike {
// 	out := make([]*model.Hike, len(bs))

// 	for i, b := range bs {
// 		out[i] = toHike(b)
// 	}
// 	return out
// }

func (*repo) JoinHike(hikeId string, token string) string {
	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return "can not find hikes collection"
	}

	log.Println("token is" + token)
	userCollection, err := db.GetDBCollection("users")
	var user *model.User
	err = userCollection.FindOne(context.TODO(), bsonmongo.D{{"token", token}}).Decode(&user)
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

	objID, _ := primitive.ObjectIDFromHex(hikeId)
	_, err2 := collection.UpdateOne(
		context.TODO(),
		bsonmongo.M{"_id": objID},
		bsonmongo.D{
			{"$push", bsonmongo.D{{"members", member}}},
		},
	)
	if err2 != nil {
		return "can not join the hiking event"
	}
	return ""
}

func (*repo) LeaveHike(hikeId string, token string) string {
	collection, err := db.GetDBCollection("hikes")
	if err != nil {
		return "can not find hikes collection"
	}

	objID, _ := primitive.ObjectIDFromHex(hikeId)
	_, err2 := collection.UpdateOne(
		context.TODO(),
		bsonmongo.M{"_id": objID},
		bsonmongo.D{
			{"$pull", bsonmongo.D{{"members", bsonmongo.M{"token": token}}}},
		},
	)
	if err2 != nil {
		return "can not leave the hiking event"
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
		bsonmongo.D{{"groupName", groupName}})
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
