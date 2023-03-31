package feature

import (
	"time"

	"github.com/turao/topics/feature-flags/entity/namespace"
	"github.com/turao/topics/metadata"
)

type config struct {
	id        ID
	namespace namespace.ID

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

func NewConfig(opts ...ConfigOption) (config, []error) {
	cfg := config{
		id:        ID("default"),
		tenancy:   metadata.TenancyTesting,
		createdAt: time.Now(),
	}

	return cfg.WithOptions(opts...)
}

func (cfg config) WithOptions(opts ...ConfigOption) (config, []error) {
	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			errs = append(errs, err)
		}
	}
	return cfg, errs
}
