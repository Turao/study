package message

import (
	"context"

	apiV1 "github.com/turao/topics/api/messages/v1"
	"github.com/turao/topics/messages/entity/message"
)

type MessageRepository interface {
	Save(ctx context.Context, message message.Message) error
}

type service struct {
	messageRepository MessageRepository
}

var _ apiV1.Messages = (*service)(nil)

func NewService(
	messageRepository MessageRepository,
) (*service, error) {
	return &service{
		messageRepository: messageRepository,
	}, nil
}

// SendMessage implements v1.Messages
func (*service) SendMessage(ctx context.Context, req apiV1.SendMessageRequest) (apiV1.SendMessageResponse, error) {
	panic("unimplemented")
}
