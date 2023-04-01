package message

import (
	"errors"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
	"github.com/turao/topics/metadata"
)

func ToModel(message message.Message) (*Model, error) {
	channels := make([]string, 0, len(message.Channels()))
	for channelID := range message.Channels() {
		channels = append(channels, channelID.String())
	}

	return &Model{
		ID:        message.ID().String(),
		Content:   message.Content(),
		Channels:  channels,
		Tenancy:   message.Tenancy().String(),
		CreatedAt: message.CreatedAt(),
		DeletedAt: message.DeletedAt(),
	}, nil
}

func ToEntity(model Model) (message.Message, error) {
	channels := make(map[channel.ID]struct{})
	for _, channelID := range model.Channels {
		channels[channel.ID(channelID)] = struct{}{}
	}

	cfg, errs := message.NewConfig(
		message.WithID(message.ID(model.ID)),
		message.WithContent(model.Content),
		message.WithChannels(channels),
		message.WithTenancy(metadata.Tenancy(model.Tenancy)),
		message.WithCreatedAt(model.CreatedAt),
		message.WithDeletedAt(model.DeletedAt),
	)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	entity := message.NewMessage(cfg)
	return entity, nil
}
