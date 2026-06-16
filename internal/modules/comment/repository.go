package comment

type CommentRepository interface {
	Create(comment Comment) error
	GetAll() ([]Comment, error)
	GetById(id int) (Comment, error)
	Update(id int, newComment Comment) error
	Delete(id int) error
}
