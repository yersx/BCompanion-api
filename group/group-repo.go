package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group, token string) string
	GetUserGroups(token string) ([]*model.Group, error)
	GetAllGroups() ([]*model.Group, error)
	GetGroup(groupName string) (*model.Group, error)
	JoinGroup(groupName string, token string) string
	LeaveGroup(groupName string, token string) string
}
