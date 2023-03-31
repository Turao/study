package channel

import (
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
	cfg config
}

var _ Channel = (*channel)(nil)

func NewChannel(cfg config) *channel {
	return &channel{
		cfg: cfg,
	}
}

func (n channel) ID() ID {
	return n.cfg.id
}

func (n channel) Tenancy() metadata.Tenancy {
	return n.cfg.tenancy
}

func (n channel) CreatedAt() time.Time {
	return n.cfg.createdAt
}

func (n channel) DeletedAt() *time.Time {
	return n.cfg.deletedAt
}
