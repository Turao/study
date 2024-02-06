package membership

import (
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/channels/entity/membership"
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

func (r *repository) FindByID(ctx context.Context, id membership.ID) (membership.Membership, error) {
	var model Model
	err := r.database.GetContext(
		ctx,
		&model,
		`SELECT channel_id, user_id, version, tenancy, created_at, deleted_at 
		FROM memberships 
		WHERE id=? 
		ORDER BY version DESC 
		LIMIT 1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}

func (r *repository) Save(ctx context.Context, membership membership.Membership) error {
	model, err := ToModel(membership)
	if err != nil {
		return err
	}

	_, err = r.database.NamedExecContext(
		ctx,
		`INSERT INTO memberships (channel_id, user_id, version, tenancy, created_at, deleted_at) 
		VALUES (:channel_id, :user_id, :version, :tenancy, :created_at, :deleted_at)`,
		model,
	)
	if err != nil {
		log.Println("error while saving", err)
		return err
	}

	return nil
}
