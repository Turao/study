package server

import (
	"context"

	apiV1 "github.com/turao/topics/messages/api/v1"
	proto "github.com/turao/topics/proto/messages"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	proto.UnimplementedMessagesServer
	service apiV1.Messages
}

var _ proto.MessagesServer = (*server)(nil)

func NewServer(service apiV1.Messages) (*server, error) {
	return &server{
		service: service,
	}, nil
}

// GetMessages implements messages.MessagesServer.
func (s *server) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.GetMessagesResponse, error) {
	res, err := s.service.GetMessages(ctx, apiV1.GetMessagesRequest{
		ChannelID: req.GetChannelId(),
	})
	if err != nil {
		return nil, err
	}

	messages := []*proto.MessageInfo{}
	for _, msg := range res.Messages {
		message := &proto.MessageInfo{
			Id:        msg.ID,
			AuthorId:  msg.Author,
			Content:   msg.Content,
			Tenancy:   msg.Tenancy,
			CreatedAt: timestamppb.New(msg.CreatedAt),
		}

		if msg.DeletedAt != nil {
			message.DeletedAt = timestamppb.New(*msg.DeletedAt)
		}

		messages = append(messages, message)
	}

	return &proto.GetMessagesResponse{
		Messages: messages,
	}, nil
}

// SendMessage implements messages.MessagesServer.
func (s *server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	_, err := s.service.SendMessage(ctx, apiV1.SendMessageRequest{
		AuthorID:  req.GetAutorId(),
		ChannelID: req.GetChannelId(),
		Content:   req.GetContent(),
		Tenancy:   req.GetTenancy(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.SendMessageResponse{}, nil
}

// StreamMessages implements messages.MessagesServer.
func (s *server) StreamMessages(req *proto.StreamMessagesRequest, stream proto.Messages_StreamMessagesServer) error {
	res, err := s.service.GetMessageStream(stream.Context(), apiV1.GetMessageStreamRequest{
		ChannelID: req.GetChannelId(),
	})
	if err != nil {
		return err
	}

	for msg := range res.Messages {
		messageInfo := &proto.MessageInfo{
			Id:        msg.ID,
			Version:   msg.Version,
			AuthorId:  msg.Author,
			Content:   msg.Content,
			Tenancy:   msg.Tenancy,
			CreatedAt: timestamppb.New(msg.CreatedAt),
		}

		if msg.DeletedAt != nil {
			messageInfo.DeletedAt = timestamppb.New(*msg.DeletedAt)
		}

		err = stream.Send(&proto.StreamMessagesResponse{
			Message: messageInfo,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
