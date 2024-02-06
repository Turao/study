package membership

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
)

type MembershipOption func(*membership) error

func WithID(id ID) MembershipOption {
	return func(ch *membership) error {

		ch.id = id
		return nil
	}
}

func WithVersion(version uint32) MembershipOption {
	return func(ch *membership) error {
		ch.version = version
		return nil
	}
}

func WithTenancy(tenancy metadata.Tenancy) MembershipOption {
	return func(ch *membership) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		ch.tenancy = tenancy
		return nil
	}
}

func WithCreatedAt(createdAt time.Time) MembershipOption {
	return func(ch *membership) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		ch.createdAt = createdAt
		return nil
	}
}

func WithDeletedAt(deletedAt *time.Time) MembershipOption {
	return func(ch *membership) error {
		if deletedAt != nil && deletedAt.Before(ch.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		ch.deletedAt = deletedAt
		return nil
	}
}
