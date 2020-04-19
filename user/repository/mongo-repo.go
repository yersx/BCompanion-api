package repository

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type repo struct{}

func NewMongoRepository() UserRepository {
	return &repo{}
}

func (*repo) SignUser(user model.User, authType string) (model.TokenResult, int) {
	var res model.TokenResult

	collection, err := db.GetDBCollection("users")
	if err != nil {
		res.Message = err.Error()
		code := 200
		return res, code
	}
	// Check if user exists in the database
	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", user.PhoneNumber}}).Decode(&user)

	if authType == "registration" {

		if err != nil {
			if err.Error() == "mongo: no documents in result" {

				// Proceed to creating user, but first generate token
				hash, err := bcrypt.GenerateFromPassword([]byte(user.PhoneNumber), 5)
				if err != nil {
					res.Message = "Error While Hashing Password, Try Again"
					code := 404
					return res, code
				}

				// store the hashed password
				user.Token = string(hash)

				// Insert User
				_, err = collection.InsertOne(context.TODO(), user)

				// Check if User Insertion Fails
				if err != nil {
					res.Message = "Error while Creating User, Try Again"
					code := 404
					return res, code
				}

				// User creation Succeeds
				res.Token = string(user.Token)
				code := 200
				return res, code
			}

			// User most likely exists
			res.Message = err.Error()
			code := 404
			return res, code
		}

		res.Message = "User already Exists!!"
		code := 404
		return res, code

	} else if authType == "login" {
		if err != nil {
			res.Message = "Not exist!"
			code := 404
			return res, code
		}

		res.Token = string(user.Token)
		code := 200
		return res, code
	}
	return res, 404
}

func (*repo) FindUser(phoneNumber string) (*model.User, error) {

	var user *model.User
	collection, err := db.GetDBCollection("users")
	if err != nil {
		return nil, err
	}

	if phoneNumber != "" {
		err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", phoneNumber}}).Decode(&user)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
