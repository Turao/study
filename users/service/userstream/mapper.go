package userstream

import (
	"errors"

	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/user"
)

func ToUserInfo(user user.User) (apiV1.UserInfo, error) {
	if user == nil {
		return apiV1.UserInfo{}, errors.New("nil user")
	}
	return apiV1.UserInfo{
		ID:        user.ID().String(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
		CreatedAt: user.CreatedAt(),
		DeletedAt: user.DeletedAt(),
	}, nil
}
