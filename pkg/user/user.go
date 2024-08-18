package user

const (
	InternalTypeAuth = "internal"
	ExternalTypeAuth = "external"
)

type User struct {
	typeAuth string
}

func New() *User {
	return &User{}
}
