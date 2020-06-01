package repository

import (
	"bcompanion/config/db"
	"bcompanion/group"
	"bcompanion/model"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type repo struct{}

func NewMongoRepository() UserRepository {
	return &repo{}
}

func (*repo) SignUser(user model.User, authType string) (string, int) {
	var res model.TokenResult

	collection, err := db.GetDBCollection("users")
	if err != nil {
		res.Message = err.Error()
		code := 404
		return "", code
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
					return "", code
				}

				// store the hashed password
				user.Token = string(hash)
				user.Photo = "https://i.ibb.co/VqncVzX/avatar.png"
				user.Status = "online"

				// Insert User
				_, err = collection.InsertOne(context.TODO(), user)

				// Check if User Insertion Fails
				if err != nil {
					res.Message = "Error while Creating User, Try Again"
					code := 404
					return "", code
				}

				// User creation Succeeds
				res.Token = string(user.Token)
				code := 200
				return res.Token, code
			}

			// User most likely exists
			res.Message = err.Error()
			code := 404
			return "", code
		}

		res.Message = "User already Exists!!"
		code := 404
		return "", code

	} else if authType == "login" {
		if err != nil {
			res.Message = "Not exist!"
			code := 404
			return "", code
		}

		res.Token = string(user.Token)
		code := 200
		return res.Token, code
	}
	return "", 404
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

func (*repo) FindToken(phoneNumber string) (*string, error) {

	type Token struct {
		Token string `bson:"token"`
	}
	var token *Token
	collection, err := db.GetDBCollection("users")
	if err != nil {
		return nil, err
	}

	type fields struct {
		Token int `bson:"token"`
	}
	projection := fields{
		Token: 1,
	}
	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", phoneNumber}}, options.FindOne().SetProjection(projection)).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token.Token, nil
}

func (*repo) FindUserProfile(phoneNumber string) (*model.UserProfile, error) {

	var up model.UserProfile
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
	up.FirstName = user.FirstName
	up.LastName = user.LastName
	up.PhoneNumber = user.PhoneNumber
	up.DateOfBirth = user.DateOfBirth
	up.City = user.City
	up.Photo = user.Photo
	up.Status = user.Status

	// upcomingHikes, err := hike.GetUpcomingByUser(user.Token)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(upcomingHikes) < 1 {
	// 	upcomingHikes = nil
	// }
	// up.UpcomingHikes = upcomingHikes

	// pastHikes, err := hike.GetPastbyUser(user.Token)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(pastHikes) < 1 {
	// 	pastHikes = nil
	// }
	// up.HikesHistory = append(upcomingHikes, pastHikes...)
	// numberOfPastHikes := len(pastHikes)
	// up.NumberOfPastHikes = strconv.Itoa(numberOfPastHikes)

	userGroups, err := group.UserGroups(user.Token)
	if err != nil {
		return nil, err
	}
	if len(userGroups) < 1 {
		userGroups = nil
		up.UpcomingHikes = nil
		up.HikesHistory = nil
		up.NumberOfPastHikes = "0"
	} else {

		past := make([]*model.Hike, 0, 100)
		upcoming := make([]*model.Hike, 0, 100)

		for _, b := range userGroups {
			upcoming = append(upcoming, b.CurrentHikes...)
			past = append(past, b.HikesHistory...)

		}
		if len(upcoming) < 1 {
			upcoming = nil
		}
		if len(past) < 1 {
			past = nil
		}
		up.HikesHistory = past
		up.UpcomingHikes = upcoming
		numberOfPastHikes := len(past)
		up.NumberOfPastHikes = strconv.Itoa(numberOfPastHikes)

	}
	numberOfGroups := len(userGroups)
	up.NumberOfGroups = strconv.Itoa(numberOfGroups)
	up.Groups = userGroups

	return &up, nil
}
