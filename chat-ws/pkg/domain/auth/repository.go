package auth

import "context"

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (User, error)
}
