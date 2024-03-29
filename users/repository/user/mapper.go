package user

import (
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/user"
)

func ToModel(user user.User) (*Model, error) {
	model := &Model{
		ID:        user.ID().String(),
		Version:   user.Version(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
		CreatedAt: user.CreatedAt(),
		DeletedAt: user.DeletedAt(),
	}

	return model, nil
}

func ToEntity(model Model) (user.User, error) {
	return user.NewUser(
		user.WithID(user.ID(model.ID)),
		user.WithVersion(model.Version),
		user.WithEmail(model.Email),
		user.WithFirstName(model.FirstName),
		user.WithLastName(model.LastName),
		user.WithTenancy(metadata.Tenancy(model.Tenancy)),
		user.WithCreatedAt(model.CreatedAt),
		user.WithDeletedAt(model.DeletedAt),
	)
}
