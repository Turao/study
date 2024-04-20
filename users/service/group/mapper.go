package group

import (
	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/group"
)

var groupMapper = GroupMapper{}

type GroupMapper struct{}

func (GroupMapper) ToGroupInfo(group group.Group) (apiV1.GroupInfo, error) {
	return apiV1.GroupInfo{
		ID:        group.ID().String(),
		Name:      group.Name(),
		Tenancy:   group.Tenancy().String(),
		CreatedAt: group.CreatedAt(),
		DeletedAt: group.DeletedAt(),
	}, nil
}
