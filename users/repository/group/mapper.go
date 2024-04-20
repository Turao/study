package group

import (
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/users/entity/group"
)

func ToModel(group group.Group) (*Model, error) {
	model := &Model{
		ID:        group.ID().String(),
		Version:   group.Version(),
		Name:      group.Name(),
		Tenancy:   group.Tenancy().String(),
		CreatedAt: group.CreatedAt(),
		DeletedAt: group.DeletedAt(),
	}

	return model, nil
}

func ToEntity(model Model) (group.Group, error) {
	return group.NewGroup(
		group.WithID(group.ID(model.ID)),
		group.WithVersion(model.Version),
		group.WithName(model.Name),
		group.WithTenancy(metadata.Tenancy(model.Tenancy)),
		group.WithCreatedAt(model.CreatedAt),
		group.WithDeletedAt(model.DeletedAt),
	)
}
