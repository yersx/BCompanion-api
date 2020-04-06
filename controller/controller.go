package controller

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/identitytoolkit/v3"
	"google.golang.org/api/option"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult

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
				w.WriteHeader(400)
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
			res.Message = string(user.Token)
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

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("captcha-token")

	// var authData model.AuthData
	phone, ok1 := r.URL.Query()["phone"]
	code := r.URL.Query()["code"]
	var res model.ResponseResult

	if !ok1 || len(phone[0]) < 1 {
		res.Message = "Url Param 'phone' is missing"
		json.NewEncoder(w).Encode(res)
		return
	}

	Phone := phone[0]
	Code := code[0]

	phoneNumber := fmt.Sprintf("+%s%s", Code, Phone)

	opt := option.WithCredentialsFile("ServiceAccountKey.json")

	log.Output(1, "phone: "+phoneNumber)
	log.Printf("token: " + tokenString)
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

	captcha, err := recaptcha.NewReCAPTCHA("6Ldg2eYUAAAAAGP8E3gTqrRQFjPFstUT4lQptSEg", recaptcha.V2, 10*time.Second)

	verErr := captcha.Verify(tokenString)
	if verErr != nil {
		log.Printf("verification error: " + err.Error())
	}
	// client, err := app.Firestore(context.Background())
	// if err != nil {
	//  log.Fatalln(err)
	// 	log.Fatalln(err)

	// }
	// defer client.Close()

	resp, err := client.Relyingparty.SendVerificationCode(&identitytoolkit.IdentitytoolkitRelyingpartySendVerificationCodeRequest{
		PhoneNumber:    phoneNumber,
		RecaptchaToken: tokenString}).Context(context.Background()).Do()
	if err != nil {
		log.Printf(err.Error())
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
	tokenString := r.Header.Get("Token")

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

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var phone model.Phone
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &phone)
	var res model.ResponseResult
	var result model.User

	if err != nil {
		res.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBCollection("users")
	if err != nil {
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(400)
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"phoneNumber", phone.PhoneNumber}}).Decode(&result)
	if err != nil {
		res.Message = "Not exist!"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = string(result.Token)
	json.NewEncoder(w).Encode(result.Token)
	return
}
