package account

import (
	"time"
)

type UserDetail struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type UserCredential struct {
	CredentialId string    `json:"credential_id" db:"credential_id"`
	UserId       string    `json:"user_id" db:"user_id"`
	Password     string    `json:"password" db:"password" validate:"min=6"`
	Role         string    `json:"role" db:"role"`
	Status       string    `json:"status" db:"status"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
	UpdatedDate  time.Time `json:"updated_date" db:"updated_date"`
}

type UserAccessInfo struct {
	CredentialId string    `json:"credential_id" db:"credential_id"`
	LastSeen     time.Time `json:"last_seen" db:"last_seen"`
}

type CreateUserParam struct {
	Username    string `json:"username" db:"username" validate:"required,min=3"`
	Password    string `json:"password" db:"password" validate:"required"`
	Role        string `db:"role"`
	Status      string
	Email       string    `json:"email" db:"email" validate:"required,email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no" validate:"required,numeric,min=10"`
	CreatedDate time.Time `db:"created_date"`
}

type UpdateUserParam struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username" validate:"required,min=3"`
	Email       string    `json:"email" db:"email" validate:"required,email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no" validate:"required,numeric,min=10"`
	UpdatedDate time.Time `db:"updated_date"`
}
