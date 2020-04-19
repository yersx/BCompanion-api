package user

import (
	"bcompanion/model"
	"bcompanion/user/repository"
)

type UserService interface {
	SignUser(user model.User, authType string) (model.TokenResult, int)
	FindUser(phoneNumber string) (*model.User, error)
}

type service struct{}

var (
	repo repository.UserRepository
)

func NewUserService(repository repository.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) SignUser(user model.User, authType string) (model.TokenResult, int) {
	return repo.SignUser(user, authType)
}

func (*service) FindUser(phoneNumber string) (*model.User, error) {
	return repo.FindUser(phoneNumber)
}
