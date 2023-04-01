package message

import (
	"errors"
	"time"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

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

func WithAuthor(author user.ID) ConfigOption {
	return func(cfg *config) error {
		if author == "" {
			return errors.New("empty author")
		}
		cfg.author = author
		return nil
	}
}

func WithChannel(channel channel.ID) ConfigOption {
	return func(cfg *config) error {
		if channel == "" {
			return errors.New("empty channel")
		}
		cfg.channel = channel
		return nil
	}
}

func WithContent(content string) ConfigOption {
	return func(cfg *config) error {
		if content == "" {
			return errors.New("empty content")
		}
		cfg.content = content
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
