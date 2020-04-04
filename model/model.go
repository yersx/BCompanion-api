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

type ResponseResult struct {
	Message string `json:"message"`
}

type AuthData struct {
	Phone     string `json:"phone"`
	Domain    string `json:"domain"`
	Code      string `json:"code"`
	Password  string `json:"password"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
}
