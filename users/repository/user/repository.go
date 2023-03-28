package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/turao/topics/users/entity/user"
)

var (
	ErrNotFound = errors.New("not found")
)

const _TableName = "users"

type repository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) (*repository, error) {
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

	_, err = r.database.QueryContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s values($1, $2, $3, $4, $5, $6, $7)", _TableName),
		model.ID,
		model.Email,
		model.FirstName,
		model.LastName,
		model.Tenancy,
		model.CreatedAt,
		model.DeletedAt,
	)

	return err
}

func (r *repository) FindByID(ctx context.Context, userID user.ID) (user.User, error) {
	var model Model
	err := r.database.QueryRowContext(
		ctx,
		fmt.Sprintf("SELECT FROM %s WHERE id = $1", _TableName),
		model.ID,
	).Scan(&model)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}
