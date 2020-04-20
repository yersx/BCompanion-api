package place

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type repo struct{}

func NewMongoRepository() PlaceRepository {
	return &repo{}
}

func (*repo) SaveCity(city model.City) error {

	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.TODO(), city)
	// Check if City Insertion Fails
	if err != nil {
		return err
	}
	return nil
}

func (*repo) GetCities() ([]*model.City, error) {

	collection, err := db.GetDBCollection("cities")
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

	return toCities(out), nil
}

func (*repo) SavePlace(place model.Place, city string) error {

	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return err
	}

	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"cityName": city},
		bson.D{
			{"$push", bson.D{{"places", place}}},
		},
	)
	if err != nil {
		return err
	}

	log.Printf("result %s", result)

	// _, err = collection.InsertOne(context.TODO(), city)
	// // Check if City Insertion Fails
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (*repo) GetPlaces(city string) ([]*model.Place, error) {

	// collection, err := db.GetDBCollection("cities")
	// if err != nil {
	// 	return nil, err
	// }

	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// defer cursor.Close(context.TODO())

	// if err != nil {
	// 	return nil, err
	// }

	// out := make([]*model.City, 0)

	// for cursor.Next(context.TODO()) {
	// 	city := new(model.City)
	// 	err := cursor.Decode(city)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	out = append(out, city)
	// }
	// if err := cursor.Err(); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func toCity(b *model.City) *model.City {
	return &model.City{
		CityName:  b.CityName,
		CityPhoto: b.CityPhoto,
	}
}

func toCities(bs []*model.City) []*model.City {
	out := make([]*model.City, len(bs))

	for i, b := range bs {
		out[i] = toCity(b)
	}
	return out
}