package v1

import "context"

type Namespaces interface {
	CreateNamespace(ctx context.Context, req CreateNamespaceRequest) (CreateNamespaceResponse, error)
	DeleteNamespace(ctx context.Context, req DeleteNamespaceRequest) (DeleteNamespaceResponse, error)
	GetNamespace(ctx context.Context, req GetNamespaceRequest) (GetNamespaceResponse, error)
}

type CreateNamespaceRequest struct{}
type CreateNamespaceResponse struct{}

type DeleteNamespaceRequest struct{}
type DeleteNamespaceResponse struct{}

type GetNamespaceRequest struct{}
type GetNamespaceResponse struct{}
