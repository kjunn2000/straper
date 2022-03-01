package admin

import "context"

type Repository interface {
	GetPaginationUsers(ctx context.Context, limit uint64, cursor string, isNext bool) ([]User, error)
	UpdateUserByAdmin(ctx context.Context, params UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}
