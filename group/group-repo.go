package group

import "bcompanion/model"

type GroupRepository interface {
	CreateGroup(group model.Group, token string) string
	GetGroups(token string) ([]*model.Group, error)
}
