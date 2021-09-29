package auth

import "time"

type User struct {
	CredentialId string    `json:"credential_id" db:"credential_id"`
	UserId       string    `json:"user_id" db:"user_id"`
	Username     string    `json:"username" db:"username"`
	Password     string    `json:"password" db:"password" validate:"min=6"`
	Role         string    `json:"role" db:"role"`
	Status       string    `json:"status" db:"status"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
	UpdatedDate  time.Time `json:"updated_date" db:"updated_date"`
}
