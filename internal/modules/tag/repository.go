package tag

type TagRepository interface {
	Create(tag Tag) error
	GetAll() ([]Tag, error)
	GetById(id int) (Tag, error)
	Update(id int, newTag Tag) error
	Delete(id int) error
	//GetByPostID(postID int) ([]Tag, error)
}
