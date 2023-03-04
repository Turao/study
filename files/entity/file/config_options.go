package file

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/movie"
)

type ConfigOption func(*config) error

func WithID(id ID) ConfigOption {
	return func(cfg *config) error {
		if id == "" {
			return errors.New("empty id")
		}
		cfg.id = id
		return nil
	}
}

func WithMovieID(movieID movie.ID) ConfigOption {
	return func(cfg *config) error {
		if movieID == "" {
			return errors.New("empty movie id")
		}
		cfg.movieID = movieID
		return nil
	}
}

func WithURI(uri string) ConfigOption {
	return func(cfg *config) error {
		if uri == "" {
			return errors.New("empty uri")
		}
		cfg.uri = uri
		return nil
	}
}

func WithSize(size int64) ConfigOption {
	return func(cfg *config) error {
		cfg.size = size
		return nil
	}
}

func WithTenancy(tenancy metadata.Tenancy) ConfigOption {
	return func(cfg *config) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		cfg.tenancy = tenancy
		return nil
	}
}

func WithCreatedAt(createdAt time.Time) ConfigOption {
	return func(cfg *config) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		cfg.createdAt = createdAt
		return nil
	}
}

func WithDeletedAt(deletedAt *time.Time) ConfigOption {
	return func(cfg *config) error {
		if deletedAt != nil && deletedAt.Before(cfg.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		cfg.deletedAt = deletedAt
		return nil
	}
}
