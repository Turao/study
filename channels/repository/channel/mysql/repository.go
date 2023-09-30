package channel

import (
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/channels/entity/channel"
)

type repository struct {
	database *sqlx.DB
}

func NewRepository(database *sqlx.DB) (*repository, error) {
	if database == nil {
		return nil, errors.New("sql database is nil")
	}
	return &repository{
		database: database,
	}, nil
}

func (r *repository) FindByID(ctx context.Context, id channel.ID) (channel.Channel, error) {
	var model Model
	err := r.database.GetContext(ctx, &model, "SELECT * FROM channels WHERE id=?", id)
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

	_, err = r.database.NamedExecContext(
		ctx,
		"REPLACE INTO channels VALUES (:id, :name, :tenancy, :created_at, :deleted_at)",
		model,
	)
	if err != nil {
		log.Println("error while saving", err)
		return err
	}

	return nil
}
