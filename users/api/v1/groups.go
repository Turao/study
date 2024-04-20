package v1

import (
	"context"
	"time"
)

type Groups interface {
	CreateGroup(ctx context.Context, req CreateGroupRequest) (CreateGroupResponse, error)
	DeleteGroup(ctx context.Context, req DeleteGroupRequest) (DeleteGroupResponse, error)
	GetGroup(ctx context.Context, req GetGroupRequest) (GetGroupResponse, error)
}

type CreateGroupRequest struct {
	Name    string `json:"name"`
	Tenancy string `json:"tenancy"`
}

type CreateGroupResponse struct {
	ID string `json:"id"`
}

type DeleteGroupRequest struct {
	ID string `json:"id"`
}

type DeleteGroupResponse struct{}

type GetGroupRequest struct {
	ID string `json:"id"`
}

type GetGroupResponse struct {
	Group GroupInfo `json:"group"`
}

type GroupInfo struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
