package channel

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
)

type ChannelOption func(*channel) error

func WithID(id ID) ChannelOption {
	return func(ch *channel) error {
		if id == "" {
			return errors.New("empty id")
		}
		ch.id = id
		return nil
	}
}

func WithVersion(version uint32) ChannelOption {
	return func(ch *channel) error {
		ch.version = version
		return nil
	}
}

func WithName(name string) ChannelOption {
	return func(ch *channel) error {
		if name == "" {
			return errors.New("empty name")
		}
		ch.name = name
		return nil
	}
}

func WithTenancy(tenancy metadata.Tenancy) ChannelOption {
	return func(ch *channel) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		ch.tenancy = tenancy
		return nil
	}
}

func WithCreatedAt(createdAt time.Time) ChannelOption {
	return func(ch *channel) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		ch.createdAt = createdAt
		return nil
	}
}

func WithDeletedAt(deletedAt *time.Time) ChannelOption {
	return func(ch *channel) error {
		if deletedAt != nil && deletedAt.Before(ch.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		ch.deletedAt = deletedAt
		return nil
	}
}
