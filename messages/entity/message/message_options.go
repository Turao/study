package message

import (
	"errors"
	"time"

	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

type MessageOption func(*message) error

func WithID(id ID) MessageOption {
	return func(msg *message) error {
		if id == "" {
			return errors.New("empty id")
		}
		msg.id = id
		return nil
	}
}

func WithVersion(version uint32) MessageOption {
	return func(m *message) error {
		m.version = version
		return nil
	}
}

func WithAuthor(author user.ID) MessageOption {
	return func(msg *message) error {
		if author == "" {
			return errors.New("empty author")
		}
		msg.author = author
		return nil
	}
}

func WithChannel(channel channel.ID) MessageOption {
	return func(msg *message) error {
		if channel == "" {
			return errors.New("empty channel")
		}
		msg.channel = channel
		return nil
	}
}

func WithContent(content string) MessageOption {
	return func(msg *message) error {
		if content == "" {
			return errors.New("empty content")
		}
		msg.content = content
		return nil
	}
}

func WithTenancy(tenancy metadata.Tenancy) MessageOption {
	return func(msg *message) error {
		if tenancy != metadata.TenancyTesting && tenancy != metadata.TenancyProduction {
			return errors.New("invalid tenancy")
		}
		msg.tenancy = tenancy
		return nil
	}
}

func WithCreatedAt(createdAt time.Time) MessageOption {
	return func(msg *message) error {
		if createdAt.After(time.Now()) {
			return errors.New("createdAt date cannot be in the future")
		}
		msg.createdAt = createdAt
		return nil
	}
}

func WithDeletedAt(deletedAt *time.Time) MessageOption {
	return func(msg *message) error {
		if deletedAt != nil && deletedAt.Before(msg.createdAt) {
			return errors.New("deletedAt date cannot be before createdAt")
		}

		msg.deletedAt = deletedAt
		return nil
	}
}
