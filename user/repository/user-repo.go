package repository

import "bcompanion/model"

type UserRepository interface {
	SignUser(user model.User, authType string) (string, int)
	FindUser(phoneNumber string) (*model.User, error)
	FindUserProfile(phoneNumber string) (*model.UserProfile, error)
	FindToken(phoneNumber string) (*string, error)
}
