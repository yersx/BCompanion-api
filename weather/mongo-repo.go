package weather

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type repo struct{}

func NewMongoRepository() WeatherRepository {
	return &repo{}
}

func (*repo) GetPlaceCoordinates(place string) (*model.Coordinate, error) {

	var coordinate *model.Coordinate

	collection, err := db.GetDBCollection("place_description")
	if err != nil {
		return nil, err
	}

	type fields struct {
		Lattitude int `bson:"latitude"`
		Longitude int `bson:"longitude"`
	}
	projection := fields{
		Lattitude: 1,
		Longitude: 1,
	}

	err = collection.FindOne(context.TODO(), bson.M{"placeName": place}, options.FindOne().SetProjection(projection)).Decode(&coordinate)
	if err != nil {
		return nil, err
		log.Println("get place points error: %v", err.Error())
	}
	return coordinate, nil
}
