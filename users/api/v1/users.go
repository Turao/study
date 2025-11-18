package v1

import (
	"context"
	"time"
)

// Users is the interface for the Users service
type Users interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) (RegisterUserResponse, error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (DeleteUserResponse, error)
	GetUserInfo(ctx context.Context, req GetUserInfoRequest) (GetUserInfoResponse, error)
}

// UsersStream is the interface for the UsersStream service
type UsersStream interface {
	StreamUsers(ctx context.Context, req StreamUsersRequest) (StreamUsersResponse, error)
}

// RegisterUserRequest is the request for the RegisterUser method
type RegisterUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}

// RegisterUserResponse is the response for the RegisterUser method
type RegisterUserResponse struct {
	ID string `json:"id"`
}

// DeleteUserRequest is the request for the DeleteUser method
type DeleteUserRequest struct {
	ID string `json:"id"`
}

// DeleteUserResponse is the response for the DeleteUser method
type DeleteUserResponse struct{}

// GetUserInfoRequest is the request for the GetUserInfo method
type GetUserInfoRequest struct {
	ID string `json:"id"`
}

// GetUserInfoResponse is the response for the GetUserInfo method
type GetUserInfoResponse struct {
	User UserInfo
}

// UserInfo is the information about a user
type UserInfo struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// StreamUsersRequest is the request for the StreamUsers method
type StreamUsersRequest struct{}

// StreamUsersResponse is the response for the StreamUsers method
type StreamUsersResponse struct {
	Users <-chan UserInfo
}
