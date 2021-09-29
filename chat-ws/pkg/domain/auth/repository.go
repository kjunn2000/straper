package auth

import "context"

type Repository interface {
	GetUserCredentialByUsername(ctx context.Context, username string) (User, error)
}
