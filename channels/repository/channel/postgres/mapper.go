package channel

import (
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
)

func ToModel(channel channel.Channel) (*Model, error) {
	return &Model{
		ID:        channel.ID().String(),
		Name:      channel.Name(),
		Tenancy:   channel.Tenancy().String(),
		CreatedAt: channel.CreatedAt(),
		DeletedAt: channel.DeletedAt(),
	}, nil
}

func ToEntity(model Model) (channel.Channel, error) {
	return channel.NewChannel(
		channel.WithID(channel.ID(model.ID)),
		channel.WithName(model.Name),
		channel.WithTenancy(metadata.Tenancy(model.Tenancy)),
		channel.WithCreatedAt(model.CreatedAt),
		channel.WithDeletedAt(model.DeletedAt),
	)
}
