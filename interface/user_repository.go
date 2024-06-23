package _interface

import (
	"context"
	"myapp/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id string) (*model.User, error) 
}