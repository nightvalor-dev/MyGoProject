package blog

type BlogService struct {
	repo BlogRepository
}

func NewBlogService(repo BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) Create(req CreateBlogRequest) error {
	blog := Blog{
		Title:   req.Title,
		Content: req.Content,
		User_id: req.UserID,
		Status:  BlogStatus(req.Status),
	}
	_, err := s.repo.Create(blog)
	return err
}

func (s *BlogService) GetAll() ([]BlogDTO, error) {
	blogs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []BlogDTO
	for _, b := range blogs {
		result = append(result, toBlogDTO(b))
	}
	return result, nil
}

func (s *BlogService) GetById(id int) (BlogDTO, error) {
	b, err := s.repo.GetById(id)
	if err != nil {
		return BlogDTO{}, err
	}
	return toBlogDTO(b), nil
}

func (s *BlogService) Update(id int, req UpdateBlogRequest) error {
	blog := Blog{
		Title:   req.Title,
		Content: req.Content,
		Status:  BlogStatus(req.Status),
	}
	return s.repo.Update(id, blog)
}

func (s *BlogService) Delete(id int) error {
	return s.repo.Delete(id)
}

func toBlogDTO(b Blog) BlogDTO {
	return BlogDTO{
		ID:      b.ID,
		Title:   b.Title,
		Content: b.Content,
		UserID:  b.User_id,
		Status:  string(b.Status),
	}
}
