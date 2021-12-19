package account

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParam) error
	GetUserDetailByUserId(ctx context.Context, userId string) (UserDetail, error)
	GetUserDetailByUsername(ctx context.Context, username string) (UserDetail, error)
	UpdateUser(ctx context.Context, user UpdateUserParam) error
	UpdateAccountPassword(ctx context.Context, userId, password string) error
	DeleteUser(ctx context.Context, userId string) error

	CreateVerifyEmailToken(ctx context.Context, token VerifyEmailToken) error
	GetVerifyEmailToken(ctx context.Context, tokenId string) (VerifyEmailToken, error)
	ValidateAccountEmail(ctx context.Context, userId, status string) error
}
