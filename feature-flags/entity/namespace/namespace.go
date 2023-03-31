package namespace

import (
	"time"

	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Namespace interface {
	ID() ID

	metadata.MultiTenant
	metadata.Auditable
}

type namespace struct {
	cfg config
}

var _ Namespace = (*namespace)(nil)

func NewNamespace(cfg config) *namespace {
	return &namespace{
		cfg: cfg,
	}
}

func (n namespace) ID() ID {
	return n.cfg.id
}

func (n namespace) Tenancy() metadata.Tenancy {
	return n.cfg.tenancy
}

func (n namespace) CreatedAt() time.Time {
	return n.cfg.createdAt
}

func (n namespace) DeletedAt() *time.Time {
	return n.cfg.deletedAt
}
