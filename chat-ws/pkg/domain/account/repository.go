package account

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParam) error
	GetUserByUserId(ctx context.Context, userId string) (User, error)
	UpdateUser(ctx context.Context, user UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}
