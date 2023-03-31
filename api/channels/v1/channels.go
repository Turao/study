package v1

import "context"

type Channels interface {
	CreateChannel(ctx context.Context, req CreateChannelRequest) (CreateChannelResponse, error)
	DeleteChannel(ctx context.Context, req DeleteChannelRequest) (DeleteChannelResponse, error)
	GetChannel(ctx context.Context, req GetChannelRequest) (GetChannelResponse, error)
}

type CreateChannelRequest struct{}
type CreateChannelResponse struct{}

type DeleteChannelRequest struct{}
type DeleteChannelResponse struct{}

type GetChannelRequest struct{}
type GetChannelResponse struct{}
