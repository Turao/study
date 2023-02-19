package user

import (
	"time"

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
	cfg config
}

var _ User = (*user)(nil)

func NewUser(cfg config) *user {
	return &user{
		cfg: cfg,
	}
}

func (u *user) ID() ID {
	return u.cfg.id
}

func (u *user) FirstName() string {
	return u.cfg.firstName
}

func (u *user) LastName() string {
	return u.cfg.lastName
}

func (u *user) Email() string {
	return u.cfg.email
}

func (u *user) Tenancy() metadata.Tenancy {
	return u.cfg.tenancy
}

func (u *user) CreatedAt() time.Time {
	return u.cfg.createdAt
}

func (u *user) DeletedAt() *time.Time {
	return u.cfg.deletedAt
}

func (u *user) Delete() {
	if u.DeletedAt() == nil {
		now := time.Now()
		u.cfg.deletedAt = &now
	}
}
