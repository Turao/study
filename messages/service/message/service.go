package message

import (
	"context"

	apiV1 "github.com/turao/topics/api/messages/v1"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type MessageRepository interface {
	Save(ctx context.Context, message message.Message) error
	ListAllByChannelID(ctx context.Context, channelID channel.ID) ([]message.Message, error)
	StreamAllByChannelID(ctx context.Context, channelID channel.ID) (<-chan message.Message, <-chan error)
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
		message.WithTenancy(metadata.Tenancy(req.Tenancy)),
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
	msgs, err := svc.messageRepository.ListAllByChannelID(ctx, channel.ID(req.ChannelID))
	if err != nil {
		return apiV1.GetMessagesResponse{}, err
	}

	msgInfos := []apiV1.MessageInfo{}
	for _, msg := range msgs {
		msgInfo, err := messageMapper.ToMessageInfo(msg)
		if err != nil {
			return apiV1.GetMessagesResponse{}, err
		}

		msgInfos = append(
			msgInfos,
			msgInfo,
		)
	}

	return apiV1.GetMessagesResponse{
		Messages: msgInfos,
	}, nil
}

// GetMessageStream implements v1.Messages.
func (svc service) GetMessageStream(ctx context.Context, req apiV1.GetMessageStreamRequest) (apiV1.GetMessageStreamResponse, error) {
	msgs, _ := svc.messageRepository.StreamAllByChannelID(ctx, channel.ID(req.ChannelID))
	// todo: consume and propagate errors

	msgInfos := make(chan apiV1.MessageInfo)
	go func() {
		defer close(msgInfos)
		for msg := range msgs {
			msgInfo, err := messageMapper.ToMessageInfo(msg)
			if err != nil {
				return
			}

			msgInfos <- msgInfo
		}
	}()

	return apiV1.GetMessageStreamResponse{
		Messages: msgInfos,
	}, nil
}
