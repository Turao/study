package v1

import (
	"context"
	"time"
)

type Users interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) (RegisterUserResponse, error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (DeleteUserResponse, error)
	GetUserInfo(ctx context.Context, req GetUserInfoRequest) (GetUserInfoResponse, error)
}

type RegisterUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}

type RegisterUserResponse struct {
	ID string `json:"id"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct{}

type GetUserInfoRequest struct {
	ID string `json:"id"`
}

type GetUserInfoResponse struct {
	User UserInfo
}

type UserInfo struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
