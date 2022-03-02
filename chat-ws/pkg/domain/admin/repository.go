package admin

import "context"

type Repository interface {
	GetUsersByCursor(ctx context.Context, limit uint64, cursor string, isNext bool) ([]User, error)
	GetUsersCount(ctx context.Context) (int, error)
	UpdateUserByAdmin(ctx context.Context, params UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}
