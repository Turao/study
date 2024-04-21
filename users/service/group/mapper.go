package group

import (
	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/group"
)

var groupMapper = GroupMapper{}

type GroupMapper struct{}

func (GroupMapper) ToGroupInfo(group group.Group) (apiV1.GroupInfo, error) {
	members := group.Members()
	memberInfos := make([]apiV1.MemberInfo, 0, len(members))
	for memberID := range members {
		memberInfos = append(
			memberInfos,
			apiV1.MemberInfo{
				ID: memberID.String(),
			},
		)
	}

	return apiV1.GroupInfo{
		ID:        group.ID().String(),
		Name:      group.Name(),
		Members:   memberInfos,
		Tenancy:   group.Tenancy().String(),
		CreatedAt: group.CreatedAt(),
		DeletedAt: group.DeletedAt(),
	}, nil
}
