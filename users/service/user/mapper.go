package user

import (
	apiV1 "github.com/turao/topics/api/users/v1"
	"github.com/turao/topics/users/entity/user"
)

var userMapper = UserMapper{}

type UserMapper struct{}

func (UserMapper) ToUserInfo(user user.User) (apiV1.UserInfo, error) {
	return apiV1.UserInfo{
		ID:        user.ID().String(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
	}, nil
}
