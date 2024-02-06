package membership

import (
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/channels/entity/membership"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

func ToModel(membership membership.Membership) (*Model, error) {
	return &Model{
		ChannelID: membership.ID().ChannelID().String(),
		UserID:    membership.ID().UserID().String(),
		Version:   membership.Version(),
		Tenancy:   membership.Tenancy().String(),
		CreatedAt: membership.CreatedAt(),
		DeletedAt: membership.DeletedAt(),
	}, nil
}

func ToEntity(model Model) (membership.Membership, error) {
	id, err := membership.NewMembershipID(
		channel.ID(model.ChannelID),
		user.ID(model.UserID),
	)
	if err != nil {
		return nil, err
	}

	return membership.NewMembership(
		id,
		membership.WithVersion(model.Version),
		membership.WithTenancy(metadata.Tenancy(model.Tenancy)),
		membership.WithCreatedAt(model.CreatedAt),
		membership.WithDeletedAt(model.DeletedAt),
	)
}
