package user

type Role string

const (
	RoleAdmin       Role = "admin"
	RoleDBA         Role = "dba"
	RolePremiumUser Role = "premium_user"
	RoleUser        Role = "user"
)
