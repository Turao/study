package feature

import (
	"time"

	"github.com/turao/topics/feature-flags/entity/namespace"
	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Feature interface {
	ID() ID
	Namespace() namespace.ID

	metadata.MultiTenant
	metadata.Auditable
}

type feature struct {
	cfg config
}

var _ Feature = (*feature)(nil)

func NewFeature(cfg config) *feature {
	return &feature{
		cfg: cfg,
	}
}

func (n feature) ID() ID {
	return n.cfg.id
}

func (n feature) Namespace() namespace.ID {
	return n.cfg.namespace
}

func (n feature) Tenancy() metadata.Tenancy {
	return n.cfg.tenancy
}

func (n feature) CreatedAt() time.Time {
	return n.cfg.createdAt
}

func (n feature) DeletedAt() *time.Time {
	return n.cfg.deletedAt
}
