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

func ToEntity(groupModel GroupModel, groupMemberModels []GroupMemberModel) (group.Group, error) {
	memberIDs := make(map[group.MemberID]struct{})
	for _, groupMemberModel := range groupMemberModels {
		if groupMemberModel.GroupID != groupMemberModel.GroupID || groupMemberModel.GroupVersion != groupModel.Version {
			continue // skip
		}
		memberID := group.MemberID(groupMemberModel.MemberID)
		memberIDs[memberID] = struct{}{}
	}

	return group.NewGroup(
		group.WithID(group.ID(groupModel.ID)),
		group.WithVersion(groupModel.Version),
		group.WithName(groupModel.Name),
		group.WithMembers(memberIDs),
		group.WithTenancy(metadata.Tenancy(groupModel.Tenancy)),
		group.WithCreatedAt(groupModel.CreatedAt),
		group.WithDeletedAt(groupModel.DeletedAt),
	)
}
