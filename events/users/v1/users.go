package v1

import "github.com/turao/topics/users/entity/user"

type UserRegistered struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}

func NewUserRegistered(user user.User) *UserRegistered {
	return &UserRegistered{
		ID:        user.ID().String(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Tenancy:   user.Tenancy().String(),
	}
}
