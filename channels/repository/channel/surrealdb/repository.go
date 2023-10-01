package channel

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/surrealdb/surrealdb.go"
	"github.com/turao/topics/channels/entity/channel"
)

type repository struct {
	database *surrealdb.DB
}

func NewRepository(database *surrealdb.DB) (*repository, error) {
	if database == nil {
		return nil, errors.New("sql database is nil")
	}
	return &repository{
		database: database,
	}, nil
}

func (r *repository) FindByID(ctx context.Context, id channel.ID) (channel.Channel, error) {
	var model Model
	funnyID := fmt.Sprintf("%s:⟨%s⟩", "channel", id)
	data, err := r.database.Select(funnyID)
	if err != nil {
		if err == surrealdb.ErrNoRow {
			log.Println("not found", funnyID, data)
			return nil, nil
		}
		log.Println("error while selecting")
		return nil, err
	}

	err = surrealdb.Unmarshal(data, &model)
	if err != nil {
		log.Println("error while unmarshaling data", data)
		return nil, err
	}

	return ToEntity(model)
}

func (r *repository) Save(ctx context.Context, channel channel.Channel) error {
	model, err := ToModel(channel)
	if err != nil {
		return err
	}

	existing, err := r.FindByID(ctx, channel.ID())
	if err != nil {
		return nil
	}

	if existing == nil {
		_, err = r.database.Create("channel", model)
	} else {
		_, err = r.database.Update("channel", model)
	}
	if err != nil {
		log.Println("error while saving", err)
		return err
	}

	return nil
}
