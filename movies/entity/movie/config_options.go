package movie

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
)

type ConfigOption func(*config) error

func Tenancy(tenancy metadata.Tenancy) ConfigOption {
	return func(cfg *config) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		cfg.tenancy = tenancy
		return nil
	}
}

func CreatedAt(createdAt time.Time) ConfigOption {
	return func(cfg *config) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		cfg.createdAt = createdAt
		return nil
	}
}

func DeletedAt(deletedAt time.Time) ConfigOption {
	return func(cfg *config) error {
		if deletedAt.Before(cfg.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}
		cfg.deletedAt = &deletedAt
		return nil
	}
}
