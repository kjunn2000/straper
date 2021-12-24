package account

import "time"

type ResetPasswordToken struct {
	TokenId     string    `json:"token_id" db:"token_id"`
	UserId      string    `json:"user_id" db:"user_id"`
	CreatedDate time.Time `json:"updated_date" db:"created_date"`
}

type UpdatePasswordParam struct {
	TokenId  string `json:"token_id" db:"token_id"`
	Password string `json:"password" db:"password" validate:"required"`
}
