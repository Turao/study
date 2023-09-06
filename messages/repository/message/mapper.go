package message

import (
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

func ToModel(message message.Message) (*Model, error) {
	return &Model{
		ID:        message.ID().String(),
		Author:    message.Author().String(),
		Channel:   message.Channel().String(),
		Content:   message.Content(),
		Tenancy:   message.Tenancy().String(),
		CreatedAt: message.CreatedAt(),
		DeletedAt: message.DeletedAt(),
	}, nil
}

func ToEntity(model Model) (message.Message, error) {
	return message.NewMessage(
		message.WithID(message.ID(model.ID)),
		message.WithAuthor(user.ID(model.Author)),
		message.WithChannel(channel.ID(model.Channel)),
		message.WithContent(model.Content),
		message.WithTenancy(metadata.Tenancy(model.Tenancy)),
		message.WithCreatedAt(model.CreatedAt),
		message.WithDeletedAt(model.DeletedAt),
	)
}
