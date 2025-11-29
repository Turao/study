package user

import (
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

// ToModel converts a User entity to a Model
func ToModel(user user.User) (*Model, error) {
	model := &Model{
		ID:        user.ID().String(),
		Version:   user.Version(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
		CreatedAt: user.CreatedAt().UnixMicro(),
	}

	if deletedAt := user.DeletedAt(); deletedAt != nil {
		ts := deletedAt.UnixMicro()
		model.DeletedAt = &ts
	}

	return model, nil
}

// ToEntity converts a Model to a User entity
func ToEntity(model Model) (user.User, error) {
	opts := []user.UserOption{
		user.WithID(user.ID(model.ID)),
		user.WithVersion(model.Version),
		user.WithEmail(model.Email),
		user.WithFirstName(model.FirstName),
		user.WithLastName(model.LastName),
		user.WithTenancy(metadata.Tenancy(model.Tenancy)),
		user.WithCreatedAt(time.UnixMicro(model.CreatedAt)),
	}

	if ts := model.DeletedAt; ts != nil {
		deletedAt := time.UnixMicro(*ts)
		opts = append(opts, user.WithDeletedAt(&deletedAt))
	}

	return user.NewUser(opts...)
}
