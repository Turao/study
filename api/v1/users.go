package v1

import "context"

type Users interface {
	RegisterUser(ctx context.Context, req RegisteUserRequest) (RegisterUserResponse, error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (DeleteUserResponse, error)
	GetUserInfo(ctx context.Context, req GetUserInfoRequest) (GetUserInfoResponse, error)
}

type RegisteUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}

type RegisterUserResponse struct {
	ID string `json:"userId"`
}

type DeleteUserRequest struct {
	ID string `json:"userId"`
}

type DeleteUserResponse struct{}

type GetUserInfoRequest struct {
	ID string `json:"userId"`
}

type GetUserInfoResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}
