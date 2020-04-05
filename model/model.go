package model

//User is A STRUCT
type User struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber,omitempty"` //PhoneNumber is a field
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	City        string `json:"city"`
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
