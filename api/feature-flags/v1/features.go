package v1

import "context"

type Features interface {
	CreateFeature(ctx context.Context, req CreateFeatureRequest) (CreateFeatureResponse, error)
	DeleteFeature(ctx context.Context, req DeleteFeatureRequest) (DeleteFeatureResponse, error)
	GetFeature(ctx context.Context, req GetFeatureRequest) (GetFeatureResponse, error)
}

type CreateFeatureRequest struct{}
type CreateFeatureResponse struct{}

type DeleteFeatureRequest struct{}
type DeleteFeatureResponse struct{}

type GetFeatureRequest struct{}
type GetFeatureResponse struct{}
