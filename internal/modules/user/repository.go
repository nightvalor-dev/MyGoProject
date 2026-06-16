package user

type UserRepository interface {
	Create(user User) error
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Update(id int, newUser User) error
	Delete(id int) error
	GetByUsername(username string) (User, error)
}
