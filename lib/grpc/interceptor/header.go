package interceptor

type Header string

const (
	HeaderTenancy Header = "x-tenancy"
	HeaderUserID  Header = "x-user-id"
)

func (h Header) String() string {
	return string(h)
}
