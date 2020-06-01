package model

//User is A STRUCT
type User struct {
	FirstName   string `json:"name" bson:"name"`
	LastName    string `json:"surname" bson:"surname"`
	Token       string `json:"-" bson:"token"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	City        string `json:"city"`
	Photo       string `json:"-" bson:"photo"`
	Status      string `json:"-" bson:"status"`
}

type Phone struct {
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber,omitempty"`
}

type TokenResult struct {
	Token   string `json:"auth_token"`
	Message string `json:"message"`
}

//AuthData is A STRUCT
type AuthData struct {
	Domain string `json:"domain"`
	Code   string `json:"code"`
	Phone  string `json:"phone"`
}

type UserProfile struct {
	FirstName         string   `json:"name"`
	LastName          string   `json:"surname"`
	PhoneNumber       string   `json:"phoneNumber"`
	DateOfBirth       string   `json:"dateOfBirth"`
	City              string   `json:"city"`
	Photo             string   `json:"photo"`
	Status            string   `json:"status"`
	UpcomingHikes     []*Hike  `json:"upcomingHikes"`
	HikesHistory      []*Hike  `json:"hikesHistory"`
	NumberOfPastHikes string   `json:"numberOfPastHikes" bson:"-"`
	Groups            []*Group `json:"groups" bson:"groups"`
	NumberOfGroups    string   `json:"numberOfGroups" bson:"-"`
}
