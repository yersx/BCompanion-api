package model

type User struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phonenumber"`
	DateOfBirth string `json:"dateOfBirth"`
	City        string `json:"city"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
