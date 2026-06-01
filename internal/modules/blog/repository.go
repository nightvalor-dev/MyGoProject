package blog

type BlogRepository interface {
	Create(blog Blog) (int, error)
	GetAll() ([]Blog, error)
	GetById(id int) (Blog, error)
	Update(id int, newBlog Blog) error
	Delete(id int) error
}
