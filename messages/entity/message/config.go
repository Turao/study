package message

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type config struct {
	id      ID
	author  user.ID
	channel channel.ID
	content string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

func NewConfig(opts ...ConfigOption) (config, []error) {
	cfg := config{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
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
