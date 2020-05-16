package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group, token string) string
	GetUserGroups(token string) ([]*model.GroupItem, error)
	GetAllGroups() ([]*model.GroupItem, error)
	GetGroup(groupName string) (*model.Group, error)
}
