package channel

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/turao/topics/channels/entity/channel"
)

type repository struct {
	database *gocql.Session
}

func NewRepository(database *gocql.Session) (*repository, error) {
	return &repository{
		database: database,
	}, nil
}

func (r *repository) FindByID(ctx context.Context, id channel.ID) (channel.Channel, error) {
	panic("not implemented yet")
}

func (r *repository) Save(ctx context.Context, channel channel.Channel) error {
	ch, err := ToModel(channel)
	if err != nil {
		return err
	}

	err = r.database.Query(
		"INSERT INTO message (id, name, tenancy, created_at, deleted_at) VALUES (?, ?, ?, ?, ?)",
		ch.ID,
		ch.Name,
		ch.Tenancy,
		ch.CreatedAt,
		ch.DeletedAt,
	).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}
