package tag

type TagService struct {
	repo TagRepository
}

func NewTagService(repo TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) Create(req CreateTagRequest) error {
	tag := Tag{
		TagName: req.TagName,
	}
	return s.repo.Create(tag)
}

func (s *TagService) GetAll() ([]TagResponse, error) {
	tags, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []TagResponse
	for _, t := range tags {
		result = append(result, toTagResponse(t))
	}
	return result, nil
}

func (s *TagService) GetById(id int) (TagResponse, error) {
	t, err := s.repo.GetById(id)
	if err != nil {
		return TagResponse{}, err
	}
	return toTagResponse(t), nil
}

func (s *TagService) Update(id int, req UpdateTagRequest) error {
	tag := Tag{
		TagName: req.TagName,
	}
	return s.repo.Update(id, tag)
}

func (s *TagService) Delete(id int) error {
	return s.repo.Delete(id)
}
