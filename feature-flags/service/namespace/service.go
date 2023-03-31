package namespace

import (
	"context"

	apiV1 "github.com/turao/topics/api/feature-flags/v1"
	"github.com/turao/topics/feature-flags/entity/namespace"
)

type NamespaceRepository interface {
	Save(ctx context.Context, namespace namespace.Namespace) error
	FindByID(ctx context.Context, namespaceID namespace.ID) (namespace.Namespace, error)
}

type service struct {
	namespaceRepository NamespaceRepository
}

var _ apiV1.Namespaces = (*service)(nil)

func NewService(
	namespaceRepository NamespaceRepository,
) (*service, error) {
	return &service{
		namespaceRepository: namespaceRepository,
	}, nil
}

// CreateNamespace implements v1.Namespaces
func (*service) CreateNamespace(ctx context.Context, req apiV1.CreateNamespaceRequest) (apiV1.CreateNamespaceResponse, error) {
	panic("unimplemented")
}

// DeleteNamespace implements v1.Namespaces
func (*service) DeleteNamespace(ctx context.Context, req apiV1.DeleteNamespaceRequest) (apiV1.DeleteNamespaceResponse, error) {
	panic("unimplemented")
}

// GetNamespace implements v1.Namespaces
func (*service) GetNamespace(ctx context.Context, req apiV1.GetNamespaceRequest) (apiV1.GetNamespaceResponse, error) {
	panic("unimplemented")
}
