package channel

import (
	"context"

	apiV1 "github.com/turao/topics/api/channels/v1"
	"github.com/turao/topics/channels/entity/channel"
)

type ChannelRepository interface {
	Save(ctx context.Context, channel channel.Channel) error
	FindByID(ctx context.Context, channelID channel.ID) (channel.Channel, error)
}

type service struct {
	channelRepository ChannelRepository
}

var _ apiV1.Channels = (*service)(nil)

func NewService(
	channelRepository ChannelRepository,
) (*service, error) {
	return &service{
		channelRepository: channelRepository,
	}, nil
}

// CreateChannel implements v1.Channels
func (*service) CreateChannel(ctx context.Context, req apiV1.CreateChannelRequest) (apiV1.CreateChannelResponse, error) {
	panic("unimplemented")
}

// DeleteChannel implements v1.Channels
func (*service) DeleteChannel(ctx context.Context, req apiV1.DeleteChannelRequest) (apiV1.DeleteChannelResponse, error) {
	panic("unimplemented")
}

// GetChannel implements v1.Channels
func (*service) GetChannelInfo(ctx context.Context, req apiV1.GetChannelInfoRequest) (apiV1.GetChannelInfoResponse, error) {
	panic("unimplemented")
}
