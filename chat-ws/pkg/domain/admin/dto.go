package admin

import "time"

type UpdateUserParam struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no"`
	Status      string    `json:"status"`
	Password    string    `json:"password" validate:"min=6"`
	UpdatedDate time.Time `json:"updated_date"`
}

type UpdateUserDetailParm struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no"`
	UpdatedDate time.Time `json:"updated_date"`
}

type UpdateCredentialParam struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Status      string    `json:"status"`
	Password    string    `json:"password" validate:"min=6"`
	UpdatedDate time.Time `json:"updated_date"`
}
