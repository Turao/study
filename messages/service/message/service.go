package message

import (
	"context"

	apiV1 "github.com/turao/topics/api/messages/v1"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
	"github.com/turao/topics/users/entity/user"
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
func (svc service) SendMessage(ctx context.Context, req apiV1.SendMessageRequest) (apiV1.SendMessageResponse, error) {
	msg, err := message.NewMessage(
		message.WithAuthor(user.ID(req.AuthorID)),
		message.WithChannel(channel.ID(req.ChannelID)),
		message.WithContent(req.Content),
	)
	if err != nil {
		return apiV1.SendMessageResponse{}, err
	}

	err = svc.messageRepository.Save(ctx, msg)
	if err != nil {
		return apiV1.SendMessageResponse{}, err
	}

	return apiV1.SendMessageResponse{}, nil
}
