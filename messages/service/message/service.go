package message

import (
	"context"
	"errors"

	apiV1 "github.com/turao/topics/api/messages/v1"
	"github.com/turao/topics/channels/entity/channel"
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
func (svc service) SendMessage(ctx context.Context, req apiV1.SendMessageRequest) (apiV1.SendMessageResponse, error) {
	// todo: method-extract api-entity mapping
	channels := make(map[channel.ID]struct{})
	for _, channelID := range req.Channels {
		channels[channel.ID(channelID)] = struct{}{}
	}

	cfg, errs := message.NewConfig(
		message.WithContent(req.Content),
		message.WithChannels(channels),
	)
	if len(errs) > 0 {
		return apiV1.SendMessageResponse{}, errors.Join(errs...)
	}
	msg := message.NewMessage(cfg)
	err := svc.messageRepository.Save(ctx, msg)
	if err != nil {
		return apiV1.SendMessageResponse{}, err
	}

	return apiV1.SendMessageResponse{}, nil
}
