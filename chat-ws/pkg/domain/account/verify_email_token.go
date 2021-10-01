package account

import "time"

type VerifyEmailToken struct {
	TokenId     string    `json:"token_id" db:"token_id"`
	UserId      string    `json:"user_id" db:"user_id"`
	CreatedDate time.Time `json:"updated_date" db:"created_date"`
}
