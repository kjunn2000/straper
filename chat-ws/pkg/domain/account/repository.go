package account

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParam) error
	GetUserDetailByUserId(ctx context.Context, userId string) (UserDetail, error)
	GetUserDetailByUsername(ctx context.Context, username string) (UserDetail, error)
	GetUserDetailByEmail(ctx context.Context, email string) (UserDetail, error)
	UpdateUser(ctx context.Context, user UpdateUserParam) error
	UpdateAccountPassword(ctx context.Context, userId, password string) error
	UpdateAccountStatus(ctx context.Context, userId, status string) error
	DeleteUser(ctx context.Context, userId string) error

	CreateVerifyEmailToken(ctx context.Context, token VerifyEmailToken) error
	GetVerifyEmailToken(ctx context.Context, tokenId string) (VerifyEmailToken, error)
	ValidateAccountEmail(ctx context.Context, userId, status string) error

	CreateResetPasswordToken(ctx context.Context, params ResetPasswordToken) error
	GetResetPasswordToken(ctx context.Context, tokenId string) (ResetPasswordToken, error)
	GetResetPasswordTokenByUserId(ctx context.Context, userId string) (ResetPasswordToken, error)
	DeleteResetPasswordToken(ctx context.Context, tokenId string) error
}
