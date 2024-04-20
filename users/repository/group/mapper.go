package group

import (
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/group"
)

func ToGroupModel(group group.Group) (*GroupModel, error) {
	groupModel := &GroupModel{
		ID:        group.ID().String(),
		Version:   group.Version(),
		Name:      group.Name(),
		Tenancy:   group.Tenancy().String(),
		CreatedAt: group.CreatedAt(),
		DeletedAt: group.DeletedAt(),
	}

	return groupModel, nil
}

func ToGroupMemberModels(group group.Group) ([]*GroupMemberModel, error) {
	groupMemberModels := []*GroupMemberModel{}
	for memberID := range group.Members() {
		groupMemberModels = append(
			groupMemberModels,
			&GroupMemberModel{
				GroupID:      group.ID().String(),
				GroupVersion: group.Version(),
				MemberID:     string(memberID),
			},
		)
	}
	return groupMemberModels, nil
}

func ToEntity(model GroupModel) (group.Group, error) {
	return group.NewGroup(
		group.WithID(group.ID(model.ID)),
		group.WithVersion(model.Version),
		group.WithName(model.Name),
		group.WithTenancy(metadata.Tenancy(model.Tenancy)),
		group.WithCreatedAt(model.CreatedAt),
		group.WithDeletedAt(model.DeletedAt),
	)
}
