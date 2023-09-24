package user

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/metadata"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type User interface {
	ID() ID

	Email() string
	FirstName() string
	LastName() string

	Delete()

	metadata.Auditable
	metadata.MultiTenant
}

type user struct {
	id ID

	email     string
	firstName string
	lastName  string

	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ User = (*user)(nil)

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

func (u *user) ID() ID {
	return u.id
}

func (u *user) FirstName() string {
	return u.firstName
}

func (u *user) LastName() string {
	return u.lastName
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Tenancy() metadata.Tenancy {
	return u.tenancy
}

func (u *user) CreatedAt() time.Time {
	return u.createdAt
}

func (u *user) DeletedAt() *time.Time {
	return u.deletedAt
}

func (u *user) Delete() {
	now := time.Now()
	u.deletedAt = &now
}
