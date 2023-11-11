package channel

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Channel interface {
	ID() ID
	Version() uint32
	Name() string

	metadata.Auditable
	metadata.Deletable
	metadata.MultiTenant
}

type channel struct {
	id      ID
	version uint32
	name    string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ Channel = (*channel)(nil)

func NewChannel(opts ...ChannelOption) (*channel, error) {
	channel := &channel{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
		version:   0,
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
func (ch channel) Version() uint32 {
	return ch.version
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

func (ch channel) Delete() {
	now := time.Now()
	ch.deletedAt = &now
	ch.version += 1
}
