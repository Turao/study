package channel

import (
	"context"
	"log"

	apiV1 "github.com/turao/topics/channels/api/v1"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/channels/entity/membership"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type ChannelRepository interface {
	Save(ctx context.Context, channel channel.Channel) error
	FindByID(ctx context.Context, channelID channel.ID) (channel.Channel, error)
}

type MembershipRepository interface {
	Save(ctx context.Context, membership membership.Membership) error
	FindByID(ctx context.Context, membershipID membership.ID) (membership.Membership, error)
}

type service struct {
	channelRepository    ChannelRepository
	membershipRepository MembershipRepository
}

var _ apiV1.Channels = (*service)(nil)

func NewService(
	channelRepository ChannelRepository,
	membershipRepository MembershipRepository,
) (*service, error) {
	return &service{
		channelRepository:    channelRepository,
		membershipRepository: membershipRepository,
	}, nil
}

// CreateChannel implements v1.Channels
func (svc service) CreateChannel(ctx context.Context, req apiV1.CreateChannelRequest) (apiV1.CreateChannelResponse, error) {
	channel, err := channel.NewChannel(
		channel.WithName(req.Name),
		channel.WithTenancy(metadata.Tenancy(req.Tenancy)),
	)
	if err != nil {
		return apiV1.CreateChannelResponse{}, err
	}

	err = svc.channelRepository.Save(ctx, channel)
	if err != nil {
		return apiV1.CreateChannelResponse{}, err
	}

	return apiV1.CreateChannelResponse{
		ID: channel.ID().String(),
	}, nil
}

// DeleteChannel implements v1.Channels
func (svc service) DeleteChannel(ctx context.Context, req apiV1.DeleteChannelRequest) (apiV1.DeleteChannelResponse, error) {
	ch, err := svc.channelRepository.FindByID(ctx, channel.ID(req.ID))
	if err != nil {
		return apiV1.DeleteChannelResponse{}, err
	}

	if ch == nil {
		return apiV1.DeleteChannelResponse{}, nil
	}

	ch.Delete()
	err = svc.channelRepository.Save(ctx, ch)
	if err != nil {
		return apiV1.DeleteChannelResponse{}, err
	}

	return apiV1.DeleteChannelResponse{}, nil
}

// GetChannel implements v1.Channels
func (svc service) GetChannelInfo(ctx context.Context, req apiV1.GetChannelInfoRequest) (apiV1.GetChannelInfoResponse, error) {
	ch, err := svc.channelRepository.FindByID(ctx, channel.ID(req.ID))
	if err != nil {
		return apiV1.GetChannelInfoResponse{}, err
	}

	chInfo, err := channelMapper.ToChannelInfo(ch)
	if err != nil {
		return apiV1.GetChannelInfoResponse{}, err
	}

	return apiV1.GetChannelInfoResponse{
		Channel: chInfo,
	}, nil
}

func (svc service) JoinChannel(ctx context.Context, req apiV1.JoinChannelRequest) (apiV1.JoinChannelResponse, error) {
	ch, err := svc.channelRepository.FindByID(ctx, channel.ID(req.ChannelID))
	if err != nil {
		return apiV1.JoinChannelResponse{}, err
	}

	membershipID, err := membership.NewMembershipID(
		channel.ID(req.ChannelID),
		user.ID(req.UserID),
	)
	if err != nil {
		return apiV1.JoinChannelResponse{}, err
	}

	var newMembership membership.Membership
	found, err := svc.membershipRepository.FindByID(ctx, membershipID)
	if err != nil {
		log.Println(err)
	}

	if found != nil {
		newMembership, err = membership.NewMembership(
			membershipID,
			membership.WithTenancy(ch.Tenancy()),
			membership.WithVersion(found.Version()+1),
		)
	} else {
		newMembership, err = membership.NewMembership(
			membershipID,
			membership.WithTenancy(ch.Tenancy()),
		)
	}

	if err != nil {
		return apiV1.JoinChannelResponse{}, err
	}

	err = svc.membershipRepository.Save(ctx, newMembership)
	if err != nil {
		return apiV1.JoinChannelResponse{}, err
	}

	return apiV1.JoinChannelResponse{}, nil
}

func (svc service) LeaveChannel(ctx context.Context, req apiV1.LeaveChannelRequest) (apiV1.LeaveChannelResponse, error) {
	membershipID, err := membership.NewMembershipID(
		channel.ID(req.ChannelID),
		user.ID(req.UserID),
	)
	if err != nil {
		return apiV1.LeaveChannelResponse{}, err
	}

	membership, err := svc.membershipRepository.FindByID(ctx, membershipID)
	if err != nil {
		return apiV1.LeaveChannelResponse{}, err
	}

	membership.Delete()

	err = svc.membershipRepository.Save(ctx, membership)
	if err != nil {
		return apiV1.LeaveChannelResponse{}, err
	}

	return apiV1.LeaveChannelResponse{}, nil
}
