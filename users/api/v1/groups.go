package v1

import (
	"context"
	"time"
)

// Groups is the interface for the Groups service
type Groups interface {
	CreateGroup(ctx context.Context, req CreateGroupRequest) (CreateGroupResponse, error)
	DeleteGroup(ctx context.Context, req DeleteGroupRequest) (DeleteGroupResponse, error)
	GetGroup(ctx context.Context, req GetGroupRequest) (GetGroupResponse, error)

	UpdateMembers(ctx context.Context, req UpdateMembersRequest) (UpdateMembersResponse, error)
	GetMemberGroups(ctx context.Context, req GetMemberGroupsRequest) (GetMemberGroupsResponse, error)
}

// CreateGroupRequest is the request for the CreateGroup method
type CreateGroupRequest struct {
	Name    string `json:"name"`
	Tenancy string `json:"tenancy"`
}

// CreateGroupResponse is the response for the CreateGroup method
type CreateGroupResponse struct {
	ID string `json:"id"`
}

// DeleteGroupRequest is the request for the DeleteGroup method
type DeleteGroupRequest struct {
	ID string `json:"id"`
}

// DeleteGroupResponse is the response for the DeleteGroup method
type DeleteGroupResponse struct{}

// GetGroupRequest is the request for the GetGroup method
type GetGroupRequest struct {
	ID string `json:"id"`
}

// GetGroupResponse is the response for the GetGroup method
type GetGroupResponse struct {
	Group GroupInfo `json:"group"`
}

// GroupInfo is the information about a group
type GroupInfo struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Members   []MemberInfo `json:"members"`
	Tenancy   string       `json:"tenancy"`
	CreatedAt time.Time    `json:"createdAt"`
	DeletedAt *time.Time   `json:"deletedAt"`
}

// MemberInfo is the information about a member
type MemberInfo struct {
	ID string `json:"id"`
}

// UpdateMembersRequest is the request for the UpdateMembers method
type UpdateMembersRequest struct {
	GroupID   string   `json:"groupId"`
	MemberIDs []string `json:"memberIds"`
}

// UpdateMembersResponse is the response for the UpdateMembers method
type UpdateMembersResponse struct{}

// GetMemberGroupsRequest is the request for the GetMemberGroups method
type GetMemberGroupsRequest struct {
	MemberID string `json:"memberId"`
}

// GetMemberGroupsResponse is the response for the GetMemberGroups method
type GetMemberGroupsResponse struct {
	MemberID string            `json:"memberId"`
	Groups   []MemberGroupInfo `json:"groups"`
}

// MemberGroupInfo is the information about a member group
type MemberGroupInfo struct {
	ID string `json:"id"`
}
