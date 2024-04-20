package group

import (
	"context"
	"log"

	"github.com/turao/topics/metadata"
	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/group"
	groupentity "github.com/turao/topics/users/entity/group"
)

// GroupRepository ...
type GroupRepository interface {
	Save(ctx context.Context, group group.Group) error
	FindByID(ctx context.Context, groupID group.ID) (group.Group, error)
}

type service struct {
	groupRepository GroupRepository
}

var _ apiV1.Groups = (*service)(nil)

// NewService ...
func NewService(
	groupRepository GroupRepository,
) (*service, error) {
	return &service{
		groupRepository: groupRepository,
	}, nil
}

// CreateGroup ...
func (svc *service) CreateGroup(ctx context.Context, req apiV1.CreateGroupRequest) (apiV1.CreateGroupResponse, error) {
	log.Println("creating group", req)
	group, err := groupentity.NewGroup(
		groupentity.WithName(req.Name),
		groupentity.WithTenancy(metadata.Tenancy(req.Tenancy)),
	)
	if err != nil {
		return apiV1.CreateGroupResponse{}, err
	}

	err = svc.groupRepository.Save(ctx, group)
	if err != nil {
		return apiV1.CreateGroupResponse{}, err
	}

	log.Println("group registered succesfully")
	return apiV1.CreateGroupResponse{
		ID: group.ID().String(),
	}, nil
}

// DeleteGroup ...
func (svc *service) DeleteGroup(ctx context.Context, req apiV1.DeleteGroupRequest) (apiV1.DeleteGroupResponse, error) {
	log.Println("deleting group", req)
	group, err := svc.groupRepository.FindByID(ctx, group.ID(req.ID))
	if err != nil {
		return apiV1.DeleteGroupResponse{}, err
	}

	group.Delete()
	err = svc.groupRepository.Save(ctx, group)
	if err != nil {
		return apiV1.DeleteGroupResponse{}, err
	}

	log.Println("group deleted succesfully")
	return apiV1.DeleteGroupResponse{}, nil
}

func (svc *service) GetGroup(ctx context.Context, req apiV1.GetGroupRequest) (apiV1.GetGroupResponse, error) {
	group, err := svc.groupRepository.FindByID(ctx, group.ID(req.ID))
	if err != nil {
		return apiV1.GetGroupResponse{}, err
	}

	groupInfo, err := groupMapper.ToGroupInfo(group)
	if err != nil {
		return apiV1.GetGroupResponse{}, nil
	}

	return apiV1.GetGroupResponse{
		Group: groupInfo,
	}, nil
}

// UpdateMembers ...
func (svc *service) UpdateMembers(ctx context.Context, req apiV1.UpdateMembersRequest) (apiV1.UpdateMembersResponse, error) {
	group, err := svc.groupRepository.FindByID(ctx, groupentity.ID(req.GroupID))
	if err != nil {
		return apiV1.UpdateMembersResponse{}, err
	}

	members := make(map[groupentity.MemberID]struct{}, len(req.MemberIDs))
	for _, memberID := range req.MemberIDs {
		members[groupentity.MemberID(memberID)] = struct{}{}
	}

	group.SetMembers(members)

	err = svc.groupRepository.Save(ctx, group)
	if err != nil {
		return apiV1.UpdateMembersResponse{}, err
	}

	return apiV1.UpdateMembersResponse{}, nil
}
