package group

import (
	"bcompanion/model"
)

type GroupService interface {
	AddGroup(group model.Group, token string) string
	GetUserGroups(token string) ([]*model.Group, error)
	GetAllGroups() ([]*model.Group, error)
}

type service struct{}

var (
	groupRepo GroupRepository
)

func NewGroupService(repository GroupRepository) GroupService {
	groupRepo = repository
	return &service{}
}

func (*service) AddGroup(group model.Group, token string) string {
	return groupRepo.CreateGroup(group, token)
}

func (*service) GetUserGroups(token string) ([]*model.Group, error) {
	return groupRepo.GetUserGroups(token)
}

func (*service) GetAllGroups() ([]*model.Group, error) {
	return groupRepo.GetAllGroups()
}
