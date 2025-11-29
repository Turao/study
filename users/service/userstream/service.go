package userstream

import (
	"context"
	"errors"

	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/user"
)

type UserStreamRepository interface {
	StreamUsers(ctx context.Context) (<-chan user.User, error)
}

type service struct {
	repository UserStreamRepository
}

var _ apiV1.UsersStream = (*service)(nil)

func New(repository UserStreamRepository) (*service, error) {
	if repository == nil {
		return nil, errors.New("repository is nil")
	}

	return &service{
		repository: repository,
	}, nil
}

func (s *service) StreamUsers(ctx context.Context, req apiV1.StreamUsersRequest) (apiV1.StreamUsersResponse, error) {
	userInfos := make(chan apiV1.UserInfo)

	users, err := s.repository.StreamUsers(ctx)
	if err != nil {
		return apiV1.StreamUsersResponse{}, err
	}

	go func() {
		defer close(userInfos)
		for {
			select {
			case user := <-users:
				userInfo, err := ToUserInfo(user)
				if err != nil {
					return
				}
				userInfos <- userInfo
			case <-ctx.Done():
				return
			}
		}
	}()

	return apiV1.StreamUsersResponse{
		Users: userInfos,
	}, nil
}
