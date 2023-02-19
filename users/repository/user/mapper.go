package user

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

func ToModel(user user.User) (*Model, error) {
	model := &Model{
		ID:        user.ID().String(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
		CreatedAt: user.CreatedAt().UnixMilli(),
	}

	if user.DeletedAt() != nil {
		ts := user.DeletedAt().UnixMilli()
		model.DeletedAt = &ts
	}

	return model, nil
}

func ToEntity(model Model) (user.User, error) {
	var deletedAt *time.Time
	if model.DeletedAt != nil {
		ts := time.UnixMilli(*model.DeletedAt)
		deletedAt = &ts
	}

	usercfg, errs := user.NewConfig(
		user.WithID(user.ID(model.ID)),
		user.WithEmail(model.Email),
		user.WithFirstName(model.FirstName),
		user.WithTenancy(metadata.Tenancy(model.Tenancy)),
		user.WithCreatedAt(time.UnixMilli(model.CreatedAt)),
		user.WithDeletedAt(deletedAt),
	)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	entity := user.NewUser(usercfg)
	return entity, nil
}
