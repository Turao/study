package user

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/metadata"
)

// ID is the type for the user ID
type ID string

// String is the string representation of the user ID
func (id ID) String() string {
	return string(id)
}

// User is the interface for the user entity
type User interface {
	ID() ID
	Version() uint32

	Email() string
	FirstName() string
	LastName() string

	metadata.Auditable
	metadata.Deletable
	metadata.MultiTenant
}

// user is the implementation of the user entity
type user struct {
	id      ID
	version uint32

	email     string
	firstName string
	lastName  string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ User = (*user)(nil)

// NewUser creates a new user entity
func NewUser(opts ...UserOption) (*user, error) {
	user := &user{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
		tenancy:   metadata.TenancyTesting,
		createdAt: time.Now(),
	}

	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(user); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return user, nil
}

// ID returns the user ID
func (u user) ID() ID {
	return u.id
}

// Version returns the user version
func (u user) Version() uint32 {
	return u.version
}

// FirstName returns the user first name
func (u user) FirstName() string {
	return u.firstName
}

// LastName returns the user last name
func (u user) LastName() string {
	return u.lastName
}

// Email returns the user email
func (u user) Email() string {
	return u.email
}

// Tenancy returns the user tenancy
func (u user) Tenancy() metadata.Tenancy {
	return u.tenancy
}

// CreatedAt returns the user created at
func (u user) CreatedAt() time.Time {
	return u.createdAt
}

// DeletedAt returns the user deleted at
func (u user) DeletedAt() *time.Time {
	return u.deletedAt
}

// Delete marks the user as deleted
func (u *user) Delete() {
	now := time.Now()
	u.deletedAt = &now
	u.version += 1
}
