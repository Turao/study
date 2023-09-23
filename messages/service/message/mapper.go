package message

import (
	apiV1 "github.com/turao/topics/messages/api/v1"
	"github.com/turao/topics/messages/entity/message"
)

var messageMapper = MessageMapper{}

type MessageMapper struct{}

func (MessageMapper) ToMessageInfo(message message.Message) (apiV1.MessageInfo, error) {
	return apiV1.MessageInfo{
		ID:        message.ID().String(),
		Author:    message.Author().String(),
		Content:   message.Content(),
		Tenancy:   message.Tenancy().String(),
		CreatedAt: message.CreatedAt(),
		DeletedAt: message.DeletedAt(),
	}, nil
}
