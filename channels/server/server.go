package server

import (
	"context"

	apiV1 "github.com/turao/topics/channels/api/v1"
	proto "github.com/turao/topics/proto/channels"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	proto.UnimplementedChannelsServer
	service apiV1.Channels
}

var _ proto.ChannelsServer = (*server)(nil)

func NewServer(service apiV1.Channels) (*server, error) {
	return &server{
		service: service,
	}, nil
}

// CreateChannel ...
func (s *server) CreateChannel(ctx context.Context, req *proto.CreateChannelRequest) (*proto.CreateChannelResponse, error) {
	res, err := s.service.CreateChannel(ctx, apiV1.CreateChannelRequest{
		Name:    req.GetName(),
		Tenancy: req.GetTenancy(),
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateChannelResponse{
		Id: res.ID,
	}, nil
}

// DeleteChannel ...
func (s *server) DeleteChannel(ctx context.Context, req *proto.DeleteChannelRequest) (*proto.DeleteChannelResponse, error) {
	_, err := s.service.DeleteChannel(ctx, apiV1.DeleteChannelRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &proto.DeleteChannelResponse{}, nil
}

// GetChannelInfo ...
func (s *server) GetChannelInfo(ctx context.Context, req *proto.GetChannelInfoRequest) (*proto.GetChannelInfoResponse, error) {
	res, err := s.service.GetChannelInfo(ctx, apiV1.GetChannelInfoRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	channelInfo := &proto.ChannelInfo{
		Id:        res.Channel.ID,
		Name:      res.Channel.Name,
		Tenancy:   res.Channel.Tenancy,
		CreatedAt: timestamppb.New(res.Channel.CreatedAt),
	}
	if res.Channel.DeletedAt != nil {
		channelInfo.DeletedAt = timestamppb.New(*res.Channel.DeletedAt)
	}

	return &proto.GetChannelInfoResponse{
		Channel: channelInfo,
	}, nil
}
