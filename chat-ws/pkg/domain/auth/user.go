package auth

type User struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password" validate:"min=6"`
	Role        string    `json:"role" db:"role"`
	Email       string    `json:"email" db:"email" validate:"email"`
	PhoneNo     string    `json:"phone_no" db:"phone_no"`
}
