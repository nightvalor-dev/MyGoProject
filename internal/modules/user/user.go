package user

import "time"

type User struct {
	ID           int       `db:"id"            json:"id"`
	Username     string    `db:"username"      json:"username"`
	Email        string    `db:"email"         json:"email"`
	Phone        string    `db:"phone"         json:"phone"`
	Bio          string    `db:"bio"           json:"bio"`
	PasswordHash string    `db:"password_hash" json:"-"`
	AssignedRole Role      `db:"role"          json:"role"`
	CreatedAt    time.Time `db:"created_at"    json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"    json:"updated_at"`
}
