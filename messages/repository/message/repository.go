package message

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/turao/topics/messages/entity/message"
)

type repository struct {
	database *gocql.Session
}

func NewRepository(database *gocql.Session) (*repository, error) {
	return &repository{
		database: database,
	}, nil
}

func (r *repository) Save(ctx context.Context, message message.Message) error {
	msg, err := ToModel(message)
	if err != nil {
		return err
	}

	err = r.database.Query(
		"INSERT INTO message (id, author, channel, content, tenancy, created_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		msg.ID,
		msg.Author,
		msg.Channel,
		msg.Content,
		msg.Tenancy,
		msg.CreatedAt,
		msg.DeletedAt,
	).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}
