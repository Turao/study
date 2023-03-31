package feature

import (
	"context"

	apiV1 "github.com/turao/topics/api/feature-flags/v1"
	"github.com/turao/topics/feature-flags/entity/feature"
)

type FeatureRepository interface {
	Save(ctx context.Context, feature feature.Feature) error
	FindByID(ctx context.Context, featureID feature.ID) (feature.Feature, error)
}

type service struct {
	featureRepository FeatureRepository
}

var _ apiV1.Features = (*service)(nil)

func NewService(
	featureRepository FeatureRepository,
) (*service, error) {
	return &service{
		featureRepository: featureRepository,
	}, nil
}

// CreateFeature implements v1.Features
func (*service) CreateFeature(ctx context.Context, req apiV1.CreateFeatureRequest) (apiV1.CreateFeatureResponse, error) {
	panic("unimplemented")
}

// DeleteFeature implements v1.Features
func (*service) DeleteFeature(ctx context.Context, req apiV1.DeleteFeatureRequest) (apiV1.DeleteFeatureResponse, error) {
	panic("unimplemented")
}

// GetFeature implements v1.Features
func (*service) GetFeature(ctx context.Context, req apiV1.GetFeatureRequest) (apiV1.GetFeatureResponse, error) {
	panic("unimplemented")
}
