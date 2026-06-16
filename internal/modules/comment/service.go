package comment

type CommentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(req CreateCommentRequest) error {
	comment := Comment{
		Content: req.Content,
		Blog_id: req.BlogID,
		User_id: req.UserID,
	}
	return s.repo.Create(comment)
}

func (s *CommentService) GetAll() ([]CommentResponse, error) {
	comments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []CommentResponse
	for _, c := range comments {
		result = append(result, toCommentResponse(c))
	}
	return result, nil
}

func (s *CommentService) GetById(id int) (CommentResponse, error) {
	c, err := s.repo.GetById(id)
	if err != nil {
		return CommentResponse{}, err
	}
	return toCommentResponse(c), nil
}

func (s *CommentService) Update(id int, req UpdateCommentRequest) error {
	comment := Comment{
		Content: req.Content,
	}
	return s.repo.Update(id, comment)
}

func (s *CommentService) Delete(id int) error {
	return s.repo.Delete(id)
}
