package group

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

type GroupOption func(*group) error

func WithID(id ID) GroupOption {
	return func(u *group) error {
		if id == "" {
			return errors.New("empty id")
		}
		u.id = id
		return nil
	}
}

func WithVersion(version uint32) GroupOption {
	return func(u *group) error {
		u.version = version
		return nil
	}
}

func WithName(name string) GroupOption {
	return func(u *group) error {
		if !isAlpha(name) {
			return errors.New("invalid first name")
		}
		u.name = name
		return nil
	}
}

func WithTenancy(tenancy metadata.Tenancy) GroupOption {
	return func(u *group) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		u.tenancy = tenancy
		return nil
	}
}

func WithCreatedAt(createdAt time.Time) GroupOption {
	return func(u *group) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		u.createdAt = createdAt
		return nil
	}
}

func WithDeletedAt(deletedAt *time.Time) GroupOption {
	return func(u *group) error {
		if deletedAt != nil && deletedAt.Before(u.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		u.deletedAt = deletedAt
		return nil
	}
}
