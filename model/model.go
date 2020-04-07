package model

//User is A STRUCT
type User struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Token       string `json:"-" bson:"token"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	City        string `json:"city"`
}

type Phone struct {
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber,omitempty"`
}

//ResponseResult is A STRUCT
type ResponseResult struct {
	Message string `json:"message"`
}

//AuthData is A STRUCT
type AuthData struct {
	Domain string `json:"domain"`
	Code   string `json:"code"`
	Phone  string `json:"phone"`
}
