package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hike struct {
	HikeID            primitive.ObjectID `json:"hikeId" bson:"_id"`
	GroupName         *string            `json:"groupName" bson:"groupName"`
	GroupPhoto        *string            `json:"groupPhoto" bson:"groupPhoto"`
	PlaceName         *string            `json:"placeName" bson:"placeName"`
	HikePhoto         *string            `json:"hikePhoto" bson:"hikePhoto"`
	HikeDescription   *string            `json:"hikeDescription" bson:"hikeDescription"`
	HikeTips          *string            `json:"hikeTips" bson:"hikeTips"`
	HikeByCar         bool               `json:"hikeByCar" bson:"hikeByCar"`
	HikePrice         *string            `json:"hikePrice" bson:"hikePrice"`
	WithOvernightStay bool               `json:"withOvernightStay" bson:"withOvernightStay"`
	StartDate         *string            `json:"startDate" bson:"startDate"`
	StartDateISO      *time.Time         `json:"-" bson:"startDateISO"`
	StartTime         *string            `json:"startTime" bson:"startTime"`
	EndDate           *string            `json:"endDate" bson:"endDate"`
	EndTime           *string            `json:"endTime" bson:"endTime"`
	GatheringCity     *string            `json:"gatheringCity" bson:"gatheringCity"`
	GatheringPlace    *string            `json:"gatheringPlace" bson:"gatheringPlace"`
	Admins            []*string          `json:"admins" bson:"admins"`
	NumberOfMembers   string             `json:"numberOfMembers" bson:"numberOfMembers"`
	Members           []*Member          `json:"members" bson:"members"`
}
