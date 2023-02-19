package user

import (
	"context"
	"errors"

	"github.com/turao/topics/users/entity/user"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	users map[string]*Model
}

func NewRepository() (*repository, error) {
	return &repository{
		users: make(map[string]*Model),
	}, nil
}

func (r *repository) Save(ctx context.Context, user user.User) error {
	model, err := ToModel(user)
	if err != nil {
		return err
	}

	r.users[user.ID().String()] = model
	return nil
}

func (r *repository) FindByID(ctx context.Context, userID user.ID) (user.User, error) {
	model, found := r.users[userID.String()]
	if !found {
		return nil, ErrNotFound
	}
	return ToEntity(*model)
}
