package metadata

import "errors"

// Tenancy is the type for the tenancy
type Tenancy string

// Tenancy constants
const (
	TenancyTesting    Tenancy = "tenancy/test"
	TenancyProduction Tenancy = "tenancy/production"
)

var (
	ErrInvalidTenancy = errors.New("invalid tenancy")
)

// MultiTenant is the interface for multi-tenant entities
type MultiTenant interface {
	Tenancy() Tenancy
}

// String is the string representation of the tenancy
func (t Tenancy) String() string {
	return string(t)
}

// NewTenancy creates a new tenancy from a string
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
