package message

import (
	"time"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Message interface {
	ID() ID
	Content() string
	Channels() map[channel.ID]struct{}

	metadata.MultiTenant
	metadata.Auditable
}

type message struct {
	cfg config
}

var _ Message = (*message)(nil)

func NewMessage(cfg config) *message {
	return &message{
		cfg: cfg,
	}
}

func (m message) ID() ID {
	return m.cfg.id
}

func (m message) Content() string {
	return m.cfg.content
}

func (m message) Channels() map[channel.ID]struct{} {
	return m.cfg.channels
}

func (m message) Tenancy() metadata.Tenancy {
	return m.cfg.tenancy
}

func (m message) CreatedAt() time.Time {
	return m.cfg.createdAt
}

func (m message) DeletedAt() *time.Time {
	return m.cfg.deletedAt
}
