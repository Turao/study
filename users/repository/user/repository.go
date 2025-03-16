package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/users/entity/user"
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

func (r *repository) Save(ctx context.Context, user user.User) error {
	model, err := ToModel(user)
	if err != nil {
		return err
	}

	_, err = r.database.NamedExecContext(
		ctx,
		`INSERT INTO users (id, version, email, first_name, last_name, tenancy, created_at, deleted_at)
		VALUES (:id, :version, :email, :first_name, :last_name, :tenancy, :created_at, :deleted_at)`,
		model,
	)

	return err
}

func (r *repository) FindByID(ctx context.Context, userID user.ID) (user.User, error) {
	var model Model
	err := r.database.GetContext(
		ctx,
		&model,
		"SELECT * FROM users WHERE id = $1 ORDER BY version DESC LIMIT 1",
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return ToEntity(model)
}
