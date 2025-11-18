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

// UserOption is the option for the user entity
type UserOption func(*user) error

func WithID(id ID) UserOption {
	return func(u *user) error {
		if id == "" {
			return errors.New("empty id")
		}
		u.id = id
		return nil
	}
}

// WithVersion sets the user version
func WithVersion(version uint32) UserOption {
	return func(u *user) error {
		u.version = version
		return nil
	}
}

// WithEmail sets the user email
func WithEmail(email string) UserOption {
	return func(u *user) error {
		if email == "" {
			return errors.New("empty email")
		}
		u.email = email
		return nil
	}
}

// WithFirstName sets the user first name
func WithFirstName(firstName string) UserOption {
	return func(u *user) error {
		if !isAlpha(firstName) {
			return errors.New("invalid first name")
		}
		u.firstName = firstName
		return nil
	}
}

// WithLastName sets the user last name
func WithLastName(lastName string) UserOption {
	return func(u *user) error {
		if !isAlpha(lastName) {
			return errors.New("invalid last name")
		}
		u.lastName = lastName
		return nil
	}
}

// WithTenancy sets the user tenancy
func WithTenancy(tenancy metadata.Tenancy) UserOption {
	return func(u *user) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		u.tenancy = tenancy
		return nil
	}
}

// WithCreatedAt sets the user created at
func WithCreatedAt(createdAt time.Time) UserOption {
	return func(u *user) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		u.createdAt = createdAt
		return nil
	}
}

// WithDeletedAt sets the user deleted at
func WithDeletedAt(deletedAt *time.Time) UserOption {
	return func(u *user) error {
		if deletedAt != nil && deletedAt.Before(u.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		u.deletedAt = deletedAt
		return nil
	}
}
