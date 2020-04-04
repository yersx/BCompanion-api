package controller

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"google.golang.org/api/identitytoolkit/v3"
	"google.golang.org/api/option"

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
		res.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBCollection("users")

	if err != nil {
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	var result model.User
	// Check if user exists in the database
	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", user.PhoneNumber}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {

			// Proceed to creating user, but first generate token
			hash, err := bcrypt.GenerateFromPassword([]byte(user.PhoneNumber), 5)

			if err != nil {
				res.Message = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}

			// store the hashed password
			user.Token = string(hash)

			// Insert User
			_, err = collection.InsertOne(context.TODO(), user)

			// Check if User Insertion Fails
			if err != nil {
				res.Message = "Error while Creating User, Try Again"
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(res)
				return
			}

			// User creation Succeeds
			res.Message = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		// User most likely exists
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}

	res.Message = "User already Exists!!"
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(res)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")

	// var user model.User
	// body, _ := ioutil.ReadAll(r.Body)
	// err := json.Unmarshal(body, &user)
	var res model.ResponseResult

	// if err != nil {
	//  res.Message = "No Fields Were Sent In"
	//  json.NewEncoder(w).Encode(res)
	//  return

	// }

	opt := option.WithCredentialsFile("ServiceAccountKey.json")

	//app, err := firebase.NewApp(context.Background(), nil, opt)
	// if err != nil {

	//  log.Fatalln(err)
	// 	log.Fatalln(err)
	// }

	client, err := identitytoolkit.NewService(context.Background(), opt)
	if err != nil {
		res.Message = "no firebase service"
		json.NewEncoder(w).Encode(res)
		return
	}

	// client, err := app.Firestore(context.Background())
	// if err != nil {
	//  log.Fatalln(err)
	// 	log.Fatalln(err)

	// }
	// defer client.Close()

	resp, err := client.Relyingparty.SendVerificationCode(&identitytoolkit.IdentitytoolkitRelyingpartySendVerificationCodeRequest{
		PhoneNumber:    "+77475652503",
		RecaptchaToken: "6LcO2rQUAAAAADaKXYb5zNNiyFEMKtayz-SgPaoY"}).Context(context.Background()).Do()
	if err != nil {
		res.Message = "can not send"
		json.NewEncoder(w).Encode(res)
		return
	}

	var idToken = resp.ServerResponse.HTTPStatusCode
	res.Message = string(idToken)
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(res)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	var result model.User
	var res model.ResponseResult

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
		//log.Fatal(err)
	}

	collection, err := db.GetDBCollection("users")

	if err != nil {
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", user.PhoneNumber}}).Decode(&result)

	if err != nil {
		res.Message = "Invalid phone"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Token), []byte(user.Token))

	if err != nil {
		res.Message = "Invalid Token"
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
		res.Message = "Error while generating token, Try Again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString

	json.NewEncoder(w).Encode(result)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	var result model.User
	var res model.ResponseResult

	collection, err := db.GetDBCollection("users")
	if err != nil {
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"token", tokenString}}).Decode(&result)
	if err != nil {
		res.Message = "Invalid token"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(result)
}
