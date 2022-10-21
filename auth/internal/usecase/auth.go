package usecase

import (
	"context"
)

// User is a user model.
type User struct {
	ID       uint64
	Name     string
	Title    string
	Gender   string
	Email    string
	Password string
	Disabled bool
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	GetByEmail(context.Context, string) (*User, error)
}
