package server

import (
	"context"

	proto "github.com/turao/topics/proto/users"
	apiV1 "github.com/turao/topics/users/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	proto.UnimplementedUsersServer
	service apiV1.Users
}

var _ proto.UsersServer = (*server)(nil)

func NewServer(service apiV1.Users) (*server, error) {
	return &server{
		service: service,
	}, nil
}

// RegisterUser ...
func (s *server) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	res, err := s.service.RegisterUser(ctx, apiV1.RegisteUserRequest{
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
	_, err := s.service.DeleteUser(ctx, apiV1.DeleteUserRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.DeleteUserResponse{}, nil
}

// GetUserInfo ...
func (s *server) GetUserInfo(ctx context.Context, req *proto.GetUserInfoRequest) (*proto.GetUserInfoResponse, error) {
	res, err := s.service.GetUserInfo(ctx, apiV1.GetUserInfoRequest{
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

	return &proto.GetUserInfoResponse{}, nil
}
