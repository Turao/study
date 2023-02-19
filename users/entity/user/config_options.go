package user

import (
	"errors"
	"regexp"
	"time"

	"github.com/turao/topics/metadata"
)

var alpha, _ = regexp.Compile("^[a-zA-Z]*$")

func isAlpha(str string) bool {
	return alpha.MatchString(str)
}

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

func WithEmail(email string) ConfigOption {
	return func(cfg *config) error {
		if email == "" {
			return errors.New("empty email")
		}
		cfg.email = email
		return nil
	}
}

func WithFirstName(firstName string) ConfigOption {
	return func(cfg *config) error {
		if !isAlpha(firstName) {
			return errors.New("invalid first name")
		}
		cfg.firstName = firstName
		return nil
	}
}

func WithLastName(lastName string) ConfigOption {
	return func(cfg *config) error {
		if !isAlpha(lastName) {
			return errors.New("invalid last name")
		}
		cfg.lastName = lastName
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
