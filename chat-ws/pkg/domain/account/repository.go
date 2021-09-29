package account

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParam) error
	GetUserDetailByUserId(ctx context.Context, userId string) (UserDetail, error)
	UpdateUser(ctx context.Context, user UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}
