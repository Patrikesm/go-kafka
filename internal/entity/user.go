package entity

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) error
	FindByName(user *User) (*User, error)
	FindId(user *User, id string) (*User, error)
}

type User struct {
	ID   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		ID:   uuid.New().String(),
		Name: name,
	}
}
