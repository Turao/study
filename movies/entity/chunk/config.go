package chunk

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/movies/entity/movie"
)

type config struct {
	id      ID
	movieID movie.ID
	uri     string
	encoded bool
	content []byte

	createdAt time.Time
	deletedAt *time.Time
}

func NewConfig(opts ...ConfigOption) (config, []error) {
	cfg := config{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
		content:   make([]byte, 0),
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
