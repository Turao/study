package v1

import "context"

type Channels interface {
	CreateChannel(ctx context.Context, req CreateChannelRequest) (CreateChannelResponse, error)
	DeleteChannel(ctx context.Context, req DeleteChannelRequest) (DeleteChannelResponse, error)
	GetChannelInfo(ctx context.Context, req GetChannelInfoRequest) (GetChannelInfoResponse, error)
}

type CreateChannelRequest struct {
	Name    string `json:"name"`
	Tenancy string `json:"tenancy"`
}

type CreateChannelResponse struct{}

type DeleteChannelRequest struct {
	ID string `json:"id"`
}

type DeleteChannelResponse struct{}

type GetChannelInfoRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Tenancy string `json:"tenancy"`
}

type GetChannelInfoResponse struct{}
