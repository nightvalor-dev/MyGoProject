package user

import "time"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
}

type UserResponse struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Bio          string    `json:"bio"`
	AssignedRole string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

func toUserResponse(u User) UserResponse {
	return UserResponse{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		Phone:        u.Phone,
		Bio:          u.Bio,
		AssignedRole: string(u.AssignedRole),
		CreatedAt:    u.CreatedAt,
	}
}
