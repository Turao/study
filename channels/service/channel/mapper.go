package channel

import (
	apiV1 "github.com/turao/topics/api/channels/v1"
	"github.com/turao/topics/channels/entity/channel"
)

var channelMapper = ChannelMapper{}

type ChannelMapper struct{}

func (ChannelMapper) ToChannelInfo(channel channel.Channel) (apiV1.ChannelInfo, error) {
	return apiV1.ChannelInfo{
		ID:        channel.ID().String(),
		Name:      channel.Name(),
		Tenancy:   channel.Tenancy().String(),
		CreatedAt: channel.CreatedAt(),
		DeletedAt: channel.DeletedAt(),
	}, nil
}
