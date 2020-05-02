package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group) error
	GetGroups(token string) ([]*model.Group, error)
}
