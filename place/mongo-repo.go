package place

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

func (*repo) GetCitiesName() ([]*string, error) {

	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return nil, err
	}
	projection := bson.D{
		{"cityName", 1},
	}

	cursor, err := collection.Find(
		context.TODO(),
		bson.D{},
		options.Find().SetProjection(projection))
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	out := make([]*model.CityName, 0)

	for cursor.Next(context.TODO()) {
		city := new(model.CityName)
		err := cursor.Decode(city)
		if err != nil {
			return nil, err
		}

		out = append(out, city)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toCitiesName(out), nil
}

func (*repo) GetCityCoordinates(city string) ([]*float64, error) {

	type Coordinate struct {
		Coordinates []*float64 `bson:"coordinates"`
	}
	var coordinate *Coordinate

	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return nil, err
	}

	type fields struct {
		Coordinate int `bson:"coordinates"`
	}
	projection := fields{
		Coordinate: 1,
	}

	err = collection.FindOne(context.TODO(), bson.M{"cityName": city}, options.FindOne().SetProjection(projection)).Decode(&coordinate)
	if err != nil {
		return nil, err
	}

	return coordinate.Coordinates, nil
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

	log.Printf("added result %v", result)

	// _, err = collection.InsertOne(context.TODO(), city)
	// // Check if City Insertion Fails
	// if err != nil {
	// 	return err
	// }
	return nil
}

type fields struct {
	ID     int `bson:"_id"`
	Places int `bson:"places"`
}

func (*repo) GetPlaces(city string) ([]*model.Place, error) {
	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return nil, err
	}

	projection := fields{
		ID:     0,
		Places: 1,
	}
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{"cityName": city},
		options.Find().SetProjection(projection))
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	var episode bson.M
	var places model.Places
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&episode); err != nil {
			return nil, err
		}
		log.Printf("found episode %v", episode)
	}
	bsonBytes, _ := bson.Marshal(episode)
	bson.Unmarshal(bsonBytes, &places)
	log.Printf("found places %v", places)

	return places.Places, nil
}

func (*repo) GetPlacesName() ([]*string, error) {
	collection, err := db.GetDBCollection("cities")
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$unwind", "$places"}},
		bson.D{{"$project", bson.D{
			{"placeName", "$places.placeName"},
			{"placePhoto", "$places.placePhoto"},
			{"cityName", "$places.cityName"},
		}}},
	}

	cursor, err := collection.Aggregate(
		context.TODO(),
		pipeline)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	out := make([]*model.Place, 0)

	for cursor.Next(context.TODO()) {
		place := new(model.Place)
		err := cursor.Decode(place)
		if err != nil {
			return nil, err
		}

		out = append(out, place)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toPlaces(out), nil
}

func toPlace(b *model.Place) *string {
	return &b.PlaceName

}

func toPlaces(bs []*model.Place) []*string {
	out := make([]*string, len(bs))

	for i, b := range bs {
		out[i] = toPlace(b)
	}
	return out
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

func toCitiesName(bs []*model.CityName) []*string {
	out := make([]*string, len(bs))

	for i, b := range bs {
		out[i] = &b.CityName
	}
	return out
}

func (*repo) SaveDescription(place model.PlaceDescription) error {

	collection, err := db.GetDBCollection("place_description")
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.TODO(), place)
	// Check if City Insertion Fails
	if err != nil {
		return err
	}
	return nil
}

func (*repo) GetDescription(placeName string) (*model.Description, error) {

	var description *model.Description
	collection, err := db.GetDBCollection("place_description")
	if err != nil {
		return nil, err
	}

	if placeName != "" {
		err = collection.FindOne(context.TODO(), bson.D{{"placeName", placeName}}).Decode(&description)
		if err != nil {
			return nil, err
		}
	}
	return description, nil
}

func (*repo) GetPlaceRoute(placeName string) (*model.PlaceRoute, error) {

	var route *model.PlaceRoute
	collection, err := db.GetDBCollection("place_description")
	if err != nil {
		return nil, err
	}

	if placeName != "" {
		err = collection.FindOne(context.TODO(), bson.D{{"placeName", placeName}}).Decode(&route)
		if err != nil {
			return nil, err
		}
	}
	return route, nil
}

func (*repo) GetPlacesRoutes(city string) ([]*model.PlaceRoute, error) {

	collection, err := db.GetDBCollection("place_description")
	if err != nil {
		log.Println("connection error!")
		return nil, err
	}
	projection := bson.D{
		{"placeName", 1},
		{"cityName", 1},
		{"placePhotos", 1},
		{"latitude", 1},
		{"longitude", 1},
		{"routeByCarText", 1},
		{"routeByWalkingText", 1},
		{"routeMap", 1},
	}

	cursor, err := collection.Find(
		context.TODO(),
		bson.M{"cityName": city},
		options.Find().SetProjection(projection))
	defer cursor.Close(context.TODO())
	if err != nil {
		log.Println("cursor error!")
		return nil, err
	}

	log.Output(1, "continue")
	out := make([]*model.PlaceRoute, 0)

	for cursor.Next(context.TODO()) {
		route := new(model.PlaceRoute)
		err := cursor.Decode(route)
		if err != nil {
			log.Println("cycle error!")
			return nil, err
		}
		out = append(out, route)
	}
	log.Println("routes: %v", out)
	if err := cursor.Err(); err != nil {
		log.Println("some error!")
		return nil, err
	}
	return toRoutes(out), nil
}

// func toRoute(b *model.PlaceRoute) *model.PlaceRoute {
// 	return b
// }

func toRoutes(bs []*model.PlaceRoute) []*model.PlaceRoute {
	out := make([]*model.PlaceRoute, len(bs))

	for i, b := range bs {
		out[i] = b
	}
	log.Println("routes in func: %v", out)
	return out
}
