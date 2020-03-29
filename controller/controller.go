// controller
package controller

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Get what was sent in
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult
	// Prepare our Error JSON Response in case there was error
	if err != nil {
		res.Error = err.Error()
		res.Result = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBCollection("users")

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	var result model.User
	// Check if user exists in the database
	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", user.PhoneNumber}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {

			// Proceed to creating user, but first generate password
			hash, err := bcrypt.GenerateFromPassword([]byte(user.PhoneNumber), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}

			// store the hashed password
			user.Token = string(hash)

			// Insert User
			_, err = collection.InsertOne(context.TODO(), user)

			// Check if User Insertion Fails
			if err != nil {
				res.Error = "Error while Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				w.WriteHeader(400)
				return
			}

			// User creation Succeeds
			res.Result = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		// User most likely exists
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}

	res.Error = "User already Exists!!"
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(400)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	var result model.User
	var res model.ResponseResult

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
		//log.Fatal(err)
	}

	collection, err := db.GetDBCollection("users")

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", user.PhoneNumber}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid phone"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Token), []byte(user.Token))

	if err != nil {
		res.Error = "Invalid Token"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstname":   result.FirstName,
		"lastname":    result.LastName,
		"phoneNumber": result.PhoneNumber,
		"dateOfBirth": result.DateOfBirth,
		"city":        result.City,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token, Try Again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString

	json.NewEncoder(w).Encode(result)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	var user model.User
	var result model.User
	var res model.ResponseResult

	collection, err := db.GetDBCollection("users")
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"token", user.Token}}).Decode(&result)
	if err != nil {
		res.Error = "Invalid token"
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(err)
}
