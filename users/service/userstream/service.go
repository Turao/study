package userstream

import (
	"context"
	"time"

	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/user"
)

type service struct{}

var _ apiV1.UsersStream = (*service)(nil)

func New() (*service, error) {
	return &service{}, nil
}

func (s *service) StreamUsers(ctx context.Context, req apiV1.StreamUsersRequest) (apiV1.StreamUsersResponse, error) {
	ticker := time.NewTicker(2 * time.Second)
	users := make(chan apiV1.UserInfo)
	go func() {
		defer close(users)
		for {
			select {
			case <-ticker.C:
				u, err := user.NewUser()
				if err != nil {
					return
				}

				userInfo, err := ToUserInfo(u)
				if err != nil {
					return
				}

				users <- userInfo
			case <-ctx.Done():
				return
			}
		}
	}()

	return apiV1.StreamUsersResponse{
		Users: users,
	}, nil
}
