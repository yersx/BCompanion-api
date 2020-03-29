package model

type User struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber,omitempty"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	City        string `json:"city"`
}

type ResponseResult struct {
	Message string `json:"message"`
}
