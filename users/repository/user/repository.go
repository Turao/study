package user

import (
	"context"
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
		`INSERT INTO users VALUES (:id, :email, :first_name, :last_name, :tenancy, :created_at, :deleted_at)
		ON CONFLICT (id) DO UPDATE SET 
		email=:email, 
		first_name=:first_name, 
		last_name=:last_name, 
		tenancy=:tenancy, 
		created_at=:created_at, 
		deleted_at=:deleted_at`,
		model,
	)

	return err
}

func (r *repository) FindByID(ctx context.Context, userID user.ID) (user.User, error) {
	var model Model
	err := r.database.GetContext(
		ctx,
		&model,
		"SELECT * FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}
