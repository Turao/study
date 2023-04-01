package message

import (
	"errors"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
	"github.com/turao/topics/metadata"
)

func ToModel(message message.Message) (*Model, error) {
	return &Model{
		ID:        message.ID().String(),
		Channel:   message.Channel().String(),
		Content:   message.Content(),
		Tenancy:   message.Tenancy().String(),
		CreatedAt: message.CreatedAt(),
		DeletedAt: message.DeletedAt(),
	}, nil
}

func ToEntity(model Model) (message.Message, error) {
	cfg, errs := message.NewConfig(
		message.WithID(message.ID(model.ID)),
		message.WithChannel(channel.ID(model.Channel)),
		message.WithContent(model.Content),
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
