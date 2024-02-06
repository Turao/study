package membership

import (
	"errors"
	"fmt"
	"time"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type ID interface {
	ChannelID() channel.ID
	UserID() user.ID
}

type id struct {
	channelID channel.ID
	userID    user.ID
}

func NewMembershipID(channelID channel.ID, userID user.ID) (ID, error) {
	if channelID == "" {
		return nil, errors.New("empty channel id")
	}
	if userID == "" {
		return nil, errors.New("empty user id")
	}
	return &id{
		channelID: channelID,
		userID:    userID,
	}, nil
}

func (id id) ChannelID() channel.ID {
	return id.channelID
}

func (id id) UserID() user.ID {
	return id.userID
}

func (id id) String() string {
	return fmt.Sprintf("%s/%s", id.channelID, id.userID)
}

type Membership interface {
	ID() ID
	Version() uint32

	metadata.Auditable
	metadata.Deletable
	metadata.MultiTenant
}

type membership struct {
	id      ID
	version uint32

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

func NewMembership(id ID, opts ...MembershipOption) (Membership, error) {
	membership := &membership{
		id:        id,
		version:   0,
		createdAt: time.Now(),
	}

	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(membership); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return membership, nil
}

func (m membership) ID() ID {
	return m.id
}

func (m membership) Version() uint32 {
	return m.version
}

func (m membership) Tenancy() metadata.Tenancy {
	return m.tenancy
}

func (m membership) CreatedAt() time.Time {
	return m.createdAt
}

func (m membership) DeletedAt() *time.Time {
	return m.deletedAt
}

func (m *membership) Delete() {
	now := time.Now()
	m.deletedAt = &now
	m.version += 1
}
