package user

import (
	"context"
	"errors"

	"github.com/turao/topics/users/entity/user"

	redis "github.com/redis/go-redis/v9"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	redis *redis.Client
}

func NewRepository(redis *redis.Client) (*repository, error) {
	return &repository{
		redis: redis,
	}, nil
}

func (r *repository) Save(ctx context.Context, user user.User) error {
	model, err := ToModel(user)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, user.ID().String(), model, 0).Err()
}

func (r *repository) FindByID(ctx context.Context, userID user.ID) (user.User, error) {
	var model Model
	err := r.redis.Get(ctx, userID.String()).Scan(&model)
	if err == redis.Nil {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}
