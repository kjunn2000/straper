package admin

import "context"

type Repository interface {
	GetUser(ctx context.Context, userId string) (User, error)
	GetUsersByCursor(ctx context.Context, param GetPaginationUsersParam) ([]User, error)
	GetUsersCount(ctx context.Context, searchStr string) (int, error)
	UpdateUserByAdmin(ctx context.Context, params UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}
