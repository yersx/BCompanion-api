package controller

import (
	"bcompanion/config/db"
	"bcompanion/model"
	"bcompanion/user"
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

type controller struct{}

var (
	userService user.UserService
)

type UserController interface {
	SignUser(w http.ResponseWriter, r *http.Request)
	FindUser(w http.ResponseWriter, r *http.Request)
	FindUserProfile(w http.ResponseWriter, r *http.Request)
	FindToken(w http.ResponseWriter, r *http.Request)
}

func NewUserController(service user.UserService) UserController {
	userService = service
	return &controller{}
}

func (*controller) SignUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	var response model.TokenResult

	err := json.Unmarshal(body, &user)
	if err != nil {
		response.Message = "No Fields Were Sent In"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}

	log.Printf("user: %s", user)

	AuthType, ok1 := r.URL.Query()["auth_type"]
	if !ok1 || len(AuthType[0]) < 1 {
		response.Message = "Url Param 'authType' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	authType := AuthType[0]

	res, code := userService.SignUser(user, authType)
	w.WriteHeader(code)
	if code == 404 {
		json.NewEncoder(w).Encode(nil)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func (*controller) FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// var res model.ResponseResult

	Phone, ok1 := r.URL.Query()["phone_number"]
	if !ok1 || len(Phone[0]) < 1 {
		// res.Message = "Url Param 'phone_number' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	phone := Phone[0]
	log.Output(1, "phone: "+phone)

	result, err := userService.FindUser(phone)
	if err != nil {
		// res.Message = err.Error()

		log.Output(1, "error 404")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	log.Printf("result %s", result)
	json.NewEncoder(w).Encode(result)
	return
}

func (*controller) FindUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// var res model.ResponseResult

	Phone, ok1 := r.URL.Query()["phone_number"]
	if !ok1 || len(Phone[0]) < 1 {
		// res.Message = "Url Param 'phone_number' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	phone := Phone[0]
	log.Output(1, "phone: "+phone)

	result, err := userService.FindUserProfile(phone)
	if err != nil {
		log.Output(1, "error 404")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}

func (*controller) FindToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// var res model.ResponseResult

	Phone, ok1 := r.URL.Query()["phone_number"]
	if !ok1 || len(Phone[0]) < 1 {
		// res.Message = "Url Param 'phone_number' is missing"
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	phone := Phone[0]
	log.Output(1, "phone: "+phone)

	result, err := userService.FindToken(phone)
	if err != nil || result == nil {
		// res.Message = err.Error()
		log.Output(1, "error 404")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(nil)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
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
		"name":        result.FirstName,
		"surname":     result.LastName,
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
	json.NewEncoder(w).Encode(res)
	return
}
