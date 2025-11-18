package userstream

import (
	apiV1 "github.com/turao/topics/users/api/v1"
	"github.com/turao/topics/users/entity/user"
)

func ToUserInfo(user user.User) (apiV1.UserInfo, error) {
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
