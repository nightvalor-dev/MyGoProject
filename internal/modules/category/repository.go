package category

type CategoryRepository interface {
	Create(category Category) error
	GetAll() ([]Category, error)
	GetById(id int) (Category, error)
	Update(id int, newCategory Category) error
	Delete(id int) error
	//GetByPostId(blogID int) (*Category, error)
}
