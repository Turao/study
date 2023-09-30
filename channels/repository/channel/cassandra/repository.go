package channel

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/turao/topics/channels/entity/channel"
)

type repository struct {
	database gocqlx.Session
}

func NewRepository(database gocqlx.Session) (*repository, error) {
	return &repository{
		database: database,
	}, nil
}

func (r *repository) FindByID(ctx context.Context, id channel.ID) (channel.Channel, error) {
	var model Model
	err := r.database.Query(_table.SelectAll()).WithContext(ctx).GetRelease(&model)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}

func (r *repository) Save(ctx context.Context, channel channel.Channel) error {
	model, err := ToModel(channel)
	if err != nil {
		return err
	}

	err = r.database.Query(_table.Insert()).WithContext(ctx).BindStruct(model).ExecRelease()
	if err != nil {
		return err
	}

	return nil
}
