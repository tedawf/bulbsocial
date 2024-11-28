package domain

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	CreatedAt  string `json:"created_at"`
	IsVerified bool   `json:"is_verified"`
	RoleID     int64  `json:"role_id"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
}
