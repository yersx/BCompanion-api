package group

import (
	"bcompanion/model"
)

type GroupService interface {
	AddGroup(group model.Group, token string) string
	GetUserGroups(token string) ([]*model.GroupItem, error)
	GetAllGroups() ([]*model.GroupItem, error)
	GetGroup(groupName string) (*model.Group, error)
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

func (*service) GetUserGroups(token string) ([]*model.GroupItem, error) {
	return groupRepo.GetUserGroups(token)
}

func (*service) GetAllGroups() ([]*model.GroupItem, error) {
	return groupRepo.GetAllGroups()
}

func (*service) GetGroup(groupName string) (*model.Group, error) {
	return groupRepo.GetGroup(groupName)
}
