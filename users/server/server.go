package server

import (
	"context"

	proto "github.com/turao/topics/proto/users"
	apiV1 "github.com/turao/topics/users/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	proto.UnimplementedUsersServer
	proto.UnimplementedGroupsServer
	userService  apiV1.Users
	groupService apiV1.Groups
}

var _ proto.UsersServer = (*server)(nil)
var _ proto.GroupsServer = (*server)(nil)

func NewServer(
	userService apiV1.Users,
	groupService apiV1.Groups,
) (*server, error) {
	return &server{
		userService:  userService,
		groupService: groupService,
	}, nil
}

// RegisterUser ...
func (s *server) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	res, err := s.userService.RegisterUser(ctx, apiV1.RegisteUserRequest{
		Email:     req.GetEmail(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Tenancy:   req.GetTenancy(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.RegisterUserResponse{
		Id: res.ID,
	}, nil
}

// DeleteUser ...
func (s *server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	_, err := s.userService.DeleteUser(ctx, apiV1.DeleteUserRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.DeleteUserResponse{}, nil
}

// GetUserInfo ...
func (s *server) GetUserInfo(ctx context.Context, req *proto.GetUserInfoRequest) (*proto.GetUserInfoResponse, error) {
	res, err := s.userService.GetUserInfo(ctx, apiV1.GetUserInfoRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	userInfo := &proto.UserInfo{
		Id:        res.User.ID,
		Email:     res.User.Email,
		FirstName: res.User.FirstName,
		LastName:  res.User.LastName,
		Tenancy:   res.User.Tenancy,
		CreatedAt: timestamppb.New(res.User.CreatedAt),
	}
	if res.User.DeletedAt != nil {
		userInfo.DeletedAt = timestamppb.New(*res.User.DeletedAt)
	}

	return &proto.GetUserInfoResponse{
		User: &proto.UserInfo{
			Id:        userInfo.Id,
			Email:     userInfo.Email,
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			Tenancy:   userInfo.Tenancy,
			CreatedAt: userInfo.CreatedAt,
			DeletedAt: userInfo.DeletedAt,
		},
	}, nil
}

// CreateGroup implements users.GroupsServer.
func (s *server) CreateGroup(ctx context.Context, req *proto.CreateGroupRequest) (*proto.CreateGroupResponse, error) {
	res, err := s.groupService.CreateGroup(ctx, apiV1.CreateGroupRequest{
		Name:    req.GetName(),
		Tenancy: req.GetTenancy(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.CreateGroupResponse{
		Id: res.ID,
	}, nil
}

// DeleteGroup implements users.GroupsServer.
func (s *server) DeleteGroup(ctx context.Context, req *proto.DeleteGroupRequest) (*proto.DeleteGroupResponse, error) {
	_, err := s.groupService.DeleteGroup(ctx, apiV1.DeleteGroupRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.DeleteGroupResponse{}, nil
}

// GetGroup implements users.GroupsServer.
func (s *server) GetGroup(ctx context.Context, req *proto.GetGroupRequest) (*proto.GetGroupResponse, error) {
	res, err := s.groupService.GetGroup(ctx, apiV1.GetGroupRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	groupInfo := &proto.GroupInfo{
		Id:        res.Group.ID,
		Name:      res.Group.Name,
		Tenancy:   res.Group.Tenancy,
		CreatedAt: timestamppb.New(res.Group.CreatedAt),
	}
	if res.Group.DeletedAt != nil {
		groupInfo.DeletedAt = timestamppb.New(*res.Group.DeletedAt)
	}

	return &proto.GetGroupResponse{
		Group: groupInfo,
	}, nil
}
