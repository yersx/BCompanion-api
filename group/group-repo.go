package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group, token string) error
	GetGroups(token string) ([]*model.Group, error)
}
