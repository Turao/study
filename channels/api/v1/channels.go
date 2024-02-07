package v1

import (
	"context"
	"time"
)

type Channels interface {
	CreateChannel(ctx context.Context, req CreateChannelRequest) (CreateChannelResponse, error)
	DeleteChannel(ctx context.Context, req DeleteChannelRequest) (DeleteChannelResponse, error)
	GetChannelInfo(ctx context.Context, req GetChannelInfoRequest) (GetChannelInfoResponse, error)

	JoinChannel(ctx context.Context, req JoinChannelRequest) (JoinChannelResponse, error)
	LeaveChannel(ctx context.Context, req LeaveChannelRequest) (LeaveChannelResponse, error)
}

type CreateChannelRequest struct {
	Name    string `json:"name"`
	Tenancy string `json:"tenancy"`
}

type CreateChannelResponse struct {
	ID string `json:"id"`
}

type DeleteChannelRequest struct {
	ID string `json:"id"`
}

type DeleteChannelResponse struct{}

type GetChannelInfoRequest struct {
	ID string `json:"id"`
}

type GetChannelInfoResponse struct {
	Channel ChannelInfo
}

type ChannelInfo struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type JoinChannelRequest struct {
	ChannelID string `json:"channelId"`
	UserID    string `json:"userId"`
}

type JoinChannelResponse struct{}

type LeaveChannelRequest struct {
	ChannelID string `json:"channelId"`
	UserID    string `json:"userId"`
}

type LeaveChannelResponse struct{}
