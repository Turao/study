package channel

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Channel interface {
	ID() ID

	metadata.MultiTenant
	metadata.Auditable
}

type channel struct {
	id ID

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ Channel = (*channel)(nil)

func NewChannel(opts ...ChannelOption) (*channel, error) {
	channel := &channel{
		id:        ID("default"),
		tenancy:   metadata.TenancyTesting,
		createdAt: time.Now(),
	}

	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(channel); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return channel, nil
}

func (n channel) ID() ID {
	return n.id
}

func (n channel) Tenancy() metadata.Tenancy {
	return n.tenancy
}

func (n channel) CreatedAt() time.Time {
	return n.createdAt
}

func (n channel) DeletedAt() *time.Time {
	return n.deletedAt
}
