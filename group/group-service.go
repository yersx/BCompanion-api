package group

import (
	"bcompanion/model"
)

type GroupService interface {
	AddGroup(group model.Group) error
	GetGroups(token string) ([]*model.Group, error)
}

type service struct{}

var (
	groupRepo GroupRepository
)

func NewGroupService(repository GroupRepository) GroupService {
	groupRepo = repository
	return &service{}
}

func (*service) AddGroup(group model.Group) error {
	return groupRepo.CreateGroup(group)
}

func (*service) GetGroups(token string) ([]*model.Group, error) {
	return groupRepo.GetGroups(token)
}
