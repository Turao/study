package group

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/turao/topics/metadata"
)

// ID is the type for the group ID
type ID string

// String is the string representation of the group ID
func (id ID) String() string {
	return string(id)
}

// MemberID is the type for the member ID
type MemberID string

// String is the string representation of the member ID
func (id MemberID) String() string {
	return string(id)
}

// Group is the interface for the group entity
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

// group is the implementation of the group entity
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

// NewGroup creates a new group entity
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

// ID returns the group ID
func (g group) ID() ID {
	return g.id
}

// Version returns the group version
func (g group) Version() uint32 {
	return g.version
}

// Name returns the group name
func (g group) Name() string {
	return g.name
}

// Members returns the group members
func (g group) Members() map[MemberID]struct{} {
	return g.members
}

// SetMembers sets the group members
func (g *group) SetMembers(members map[MemberID]struct{}) {
	g.members = members
	g.version += 1
}

// Tenancy returns the group tenancy
func (g group) Tenancy() metadata.Tenancy {
	return g.tenancy
}

// CreatedAt returns the group created at
func (g group) CreatedAt() time.Time {
	return g.createdAt
}

// DeletedAt returns the group deleted at
func (g group) DeletedAt() *time.Time {
	return g.deletedAt
}

// Delete marks the group as deleted
func (g *group) Delete() {
	now := time.Now()
	g.deletedAt = &now
	g.version += 1
}
