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
	ListAllByChannelID(ctx context.Context, channelID channel.ID) ([]message.Message, error)
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

func (svc service) GetMessages(ctx context.Context, req apiV1.GetMessagesRequest) (apiV1.GetMessagesResponse, error) {
	messages, err := svc.messageRepository.ListAllByChannelID(ctx, channel.ID(req.ChannelID))
	if err != nil {
		return apiV1.GetMessagesResponse{}, err
	}

	msgs := []apiV1.MessageInfo{}
	for _, message := range messages {
		msgs = append(
			msgs,
			apiV1.MessageInfo{
				ID:        message.ID().String(),
				Author:    message.Author().String(),
				Content:   message.Content(),
				Tenancy:   message.Tenancy().String(),
				CreatedAt: message.CreatedAt(),
				DeletedAt: message.DeletedAt(),
			},
		)
	}

	return apiV1.GetMessagesResponse{
		Messages: msgs,
	}, nil
}
