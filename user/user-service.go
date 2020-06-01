package user

import (
	"bcompanion/model"
	"bcompanion/user/repository"
)

type UserService interface {
	SignUser(user model.User, authType string) (string, int)
	FindUser(phoneNumber string) (*model.User, error)
	FindUserProfile(phoneNumber string) (*model.UserProfile, error)
	FindToken(phoneNumber string) (*string, error)
}

type service struct{}

var (
	repo repository.UserRepository
)

func NewUserService(repository repository.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) SignUser(user model.User, authType string) (string, int) {
	return repo.SignUser(user, authType)
}

func (*service) FindUser(phoneNumber string) (*model.User, error) {
	return repo.FindUser(phoneNumber)
}

func (*service) FindUserProfile(phoneNumber string) (*model.UserProfile, error) {
	return repo.FindUserProfile(phoneNumber)
}

func (*service) FindToken(phoneNumber string) (*string, error) {
	return repo.FindToken(phoneNumber)
}
