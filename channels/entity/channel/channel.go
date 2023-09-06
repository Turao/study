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
	Name() string

	metadata.MultiTenant
	metadata.Auditable
}

type channel struct {
	id   ID
	name string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ Channel = (*channel)(nil)

func NewChannel(opts ...ChannelOption) (*channel, error) {
	channel := &channel{
		id:        ID("default"),
		name:      "",
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

func (ch channel) ID() ID {
	return ch.id
}

func (ch channel) Name() string {
	return ch.name
}

func (ch channel) Tenancy() metadata.Tenancy {
	return ch.tenancy
}

func (ch channel) CreatedAt() time.Time {
	return ch.createdAt
}

func (ch channel) DeletedAt() *time.Time {
	return ch.deletedAt
}
