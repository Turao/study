package group

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/metadata"
)

type ID string
type MemberID string

func (id ID) String() string {
	return string(id)
}

type Group interface {
	ID() ID
	Version() uint32

	Name() string
	Members() map[MemberID]struct{}
	SetMembers(members map[MemberID]struct{})
	metadata.Auditable
	metadata.Deletable
	metadata.MultiTenant
}

type group struct {
	id        ID
	version   uint32
	name      string
	members   map[MemberID]struct{}
	tenancy   metadata.Tenancy
	createdAt time.Time
	deletedAt *time.Time
}

var _ Group = (*group)(nil)

func NewGroup(opts ...GroupOption) (*group, error) {
	group := &group{
		id:        ID(uuid.Must(uuid.NewV4()).String()),
		tenancy:   metadata.TenancyTesting,
		createdAt: time.Now(),
		members:   make(map[MemberID]struct{}),
	}

	errs := make([]error, 0)
	for _, opt := range opts {
		if err := opt(group); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return group, nil
}

func (g group) ID() ID {
	return g.id
}

func (g group) Version() uint32 {
	return g.version
}

func (g group) Name() string {
	return g.name
}

func (g group) Members() map[MemberID]struct{} {
	return g.members
}

func (g *group) SetMembers(members map[MemberID]struct{}) {
	g.members = members
	g.version += 1
}

func (g group) Tenancy() metadata.Tenancy {
	return g.tenancy
}

func (g group) CreatedAt() time.Time {
	return g.createdAt
}

func (g group) DeletedAt() *time.Time {
	return g.deletedAt
}

func (g *group) Delete() {
	now := time.Now()
	g.deletedAt = &now
	g.version += 1
}
