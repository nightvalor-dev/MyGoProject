package category

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(req CreateCategoryRequest) error {
	category := Category{
		CategoryName: req.CategoryName,
		Description:  req.Description,
	}
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]CategoryResponse, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []CategoryResponse
	for _, c := range categories {
		result = append(result, toCategoryResponse(c))
	}
	return result, nil
}

func (s *CategoryService) GetById(id int) (CategoryResponse, error) {
	c, err := s.repo.GetById(id)
	if err != nil {
		return CategoryResponse{}, err
	}
	return toCategoryResponse(c), nil
}

func (s *CategoryService) Update(id int, req UpdateCategoryRequest) error {
	category := Category{
		CategoryName: req.CategoryName,
		Description:  req.Description,
	}
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
