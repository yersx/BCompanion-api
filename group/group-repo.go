package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group) string
	GetGroups(token string) ([]*model.Group, error)
}
