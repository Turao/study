package metadata

type Tenancy string

const (
	TenancyTesting    Tenancy = "tenancy/test"
	TenancyProduction Tenancy = "tenancy/production"
)

type MultiTenant interface {
	Tenancy() Tenancy
}

func (t Tenancy) String() string {
	return string(t)
}
