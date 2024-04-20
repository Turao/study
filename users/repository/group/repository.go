package group

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/users/entity/group"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	database *sqlx.DB
}

func NewRepository(database *sqlx.DB) (*repository, error) {
	if database == nil {
		return nil, errors.New("database connection is nil")
	}

	return &repository{
		database: database,
	}, nil
}

func (r *repository) Save(ctx context.Context, group group.Group) error {
	model, err := ToModel(group)
	if err != nil {
		return err
	}

	_, err = r.database.NamedExecContext(
		ctx,
		`INSERT INTO groups (id, version, name, tenancy, created_at, deleted_at)
		VALUES (:id, :version, :name, :tenancy, :created_at, :deleted_at)`,
		model,
	)

	return err
}

func (r *repository) FindByID(ctx context.Context, groupID group.ID) (group.Group, error) {
	var model Model
	err := r.database.GetContext(
		ctx,
		&model,
		"SELECT * FROM groups WHERE id = $1 ORDER BY version DESC LIMIT 1",
		groupID,
	)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}
