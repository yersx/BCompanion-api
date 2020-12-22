// db
package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBCollection(name string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://mongouser:user@mongocluster-fvpsj.mongodb.net/test?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("BCompanion").Collection(name)
	return collection, nil
}
