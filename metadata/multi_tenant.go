package metadata

import "errors"

type Tenancy string

const (
	TenancyTesting    Tenancy = "tenancy/test"
	TenancyProduction Tenancy = "tenancy/production"
)

var (
	ErrInvalidTenancy = errors.New("invalid tenancy")
)

type MultiTenant interface {
	Tenancy() Tenancy
}

func (t Tenancy) String() string {
	return string(t)
}

func NewTenancy(tenancy string) (Tenancy, error) {
	switch tenancy {
	case TenancyTesting.String():
		return TenancyTesting, nil
	case TenancyTesting.String():
		return TenancyTesting, nil
	default:
		return TenancyTesting, ErrInvalidTenancy
	}
}
