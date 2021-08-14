package account

import (
	"time"
)

type User struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password" validate:"min=6"`
	Role        string    `json:"role" db:"role"`
	Email       string    `json:"email" db:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

type CreateUserParam struct {
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password" validate:"min=6"`
	Role        string    `db:"role"`
	Email       string    `json:"email" db:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no"`
	CreatedDate time.Time `db:"created_date"`
}

type UpdateUserParam struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
}
