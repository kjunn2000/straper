package admin

import "time"

type User struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no"`
	Status      string    `json:"status" db:"status"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type UserInfo struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
}
