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

func (n message) ID() ID {
	return n.cfg.id
}

func (n message) Channels() map[channel.ID]struct{} {
	return n.cfg.channels
}

func (n message) Tenancy() metadata.Tenancy {
	return n.cfg.tenancy
}

func (n message) CreatedAt() time.Time {
	return n.cfg.createdAt
}

func (n message) DeletedAt() *time.Time {
	return n.cfg.deletedAt
}
