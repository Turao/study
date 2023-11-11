package message

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Message interface {
	ID() ID
	Version() uint32

	Author() user.ID
	Channel() channel.ID
	Content() string

	metadata.Auditable
	metadata.Deletable
	metadata.MultiTenant
}

type message struct {
	id      ID
	version uint32

	author  user.ID
	channel channel.ID
	content string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ Message = (*message)(nil)

func NewMessage(opts ...MessageOption) (*message, error) {
	message := &message{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
		tenancy:   metadata.TenancyTesting,
		createdAt: time.Now(),
	}

	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(message); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return message, nil
}

func (m message) ID() ID {
	return m.id
}

func (m message) Version() uint32 {
	return m.version
}

func (m message) Author() user.ID {
	return m.author
}

func (m message) Channel() channel.ID {
	return m.channel
}

func (m message) Content() string {
	return m.content
}

func (m message) Tenancy() metadata.Tenancy {
	return m.tenancy
}

func (m message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *message) Delete() {
	now := time.Now()
	m.deletedAt = &now
	m.version += 1
}

func (m message) DeletedAt() *time.Time {
	return m.deletedAt
}
