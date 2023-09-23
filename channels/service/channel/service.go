package channel

import (
	"context"

	apiV1 "github.com/turao/topics/channels/api/v1"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
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
func (svc service) CreateChannel(ctx context.Context, req apiV1.CreateChannelRequest) (apiV1.CreateChannelResponse, error) {
	channel, err := channel.NewChannel(
		channel.WithName(req.Name),
		channel.WithTenancy(metadata.Tenancy(req.Tenancy)),
	)
	if err != nil {
		return apiV1.CreateChannelResponse{}, err
	}

	err = svc.channelRepository.Save(ctx, channel)
	if err != nil {
		return apiV1.CreateChannelResponse{}, err
	}

	return apiV1.CreateChannelResponse{}, nil
}

// DeleteChannel implements v1.Channels
func (svc service) DeleteChannel(ctx context.Context, req apiV1.DeleteChannelRequest) (apiV1.DeleteChannelResponse, error) {
	ch, err := svc.channelRepository.FindByID(ctx, channel.ID(req.ID))
	if err != nil {
		return apiV1.DeleteChannelResponse{}, err
	}

	ch = ch.Delete()
	err = svc.channelRepository.Save(ctx, ch)
	if err != nil {
		return apiV1.DeleteChannelResponse{}, err
	}

	return apiV1.DeleteChannelResponse{}, nil
}

// GetChannel implements v1.Channels
func (svc service) GetChannelInfo(ctx context.Context, req apiV1.GetChannelInfoRequest) (apiV1.GetChannelInfoResponse, error) {
	ch, err := svc.channelRepository.FindByID(ctx, channel.ID(req.ID))
	if err != nil {
		return apiV1.GetChannelInfoResponse{}, err
	}

	chInfo, err := channelMapper.ToChannelInfo(ch)
	if err != nil {
		return apiV1.GetChannelInfoResponse{}, err
	}

	return apiV1.GetChannelInfoResponse{
		Channel: chInfo,
	}, nil
}
