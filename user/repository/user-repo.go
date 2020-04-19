package repository

import "bcompanion/model"

type UserRepository interface {
	SignUser(user model.User, authType string) (model.TokenResult, int)
	FindUser(phoneNumber string) (*model.User, error)
}
