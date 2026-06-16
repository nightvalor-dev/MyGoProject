package blog

import (
	"Project2-v3/internal/modules/category"
	"Project2-v3/internal/modules/tag"
)

type BlogService struct {
	blogrepo     BlogRepository
	categoryRepo category.CategoryRepository
	tagRepo      tag.TagRepository
}

func NewBlogService(
	blogrepo BlogRepository,
	categoryRepo category.CategoryRepository,
	tagRepo tag.TagRepository,
) *BlogService {
	return &BlogService{
		blogrepo:     blogrepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (s *BlogService) Create(req CreateBlogRequest) error {

	blog := Blog{
		Title:      req.Title,
		Content:    req.Content,
		UserId:     req.UserId,
		Status:     BlogStatus(req.Status),
		CategoryId: req.CategoryId,
	}

	blogID, err := s.blogrepo.Create(blog)
	if err != nil {
		return err
	}

	for _, t := range req.Tags {

		tagID, err := s.tagRepo.GetOrCreate(t.TagName)
		if err != nil {
			return err
		}

		err = s.blogrepo.AddTag(blogID, tagID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BlogService) GetAll() ([]BlogResponse, error) {
	blogs, err := s.blogrepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []BlogResponse
	for _, b := range blogs {
		response, err := s.buildResponse(b)
		if err != nil {
			return nil, err
		}
		result = append(result, response)
	}
	return result, nil
}

func (s *BlogService) GetById(id int) (BlogResponse, error) {
	b, err := s.blogrepo.GetById(id)
	if err != nil {
		return BlogResponse{}, err
	}
	return s.buildResponse(b)
}

func (s *BlogService) Update(id int, req UpdateBlogRequest) error {
	blog := Blog{
		Title:   req.Title,
		Content: req.Content,
		Status:  BlogStatus(req.Status),
	}
	return s.blogrepo.Update(id, blog)
}

func (s *BlogService) Delete(id int) error {
	return s.blogrepo.Delete(id)
}

func (s *BlogService) buildResponse(b Blog) (BlogResponse, error) {
	cat, err := s.categoryRepo.GetByPostId(b.ID)
	if err != nil {
		return BlogResponse{}, err
	}

	tags, err := s.tagRepo.GetByPostID(b.ID)
	if err != nil {
		return BlogResponse{}, err
	}

	tagSummaries := make([]TagLocal, len(tags))
	for i, t := range tags {
		tagSummaries[i] = TagLocal{ID: t.ID, TagName: t.TagName}
	}

	return BlogResponse{
		ID:      b.ID,
		Title:   b.Title,
		Content: b.Content,
		UserID:  b.User_id,
		Status:  string(b.Status),
		Category: CategoryLocal{
			ID:           cat.ID,
			CategoryName: cat.CategoryName,
			Description:  cat.Description,
		},
		Tags: tagSummaries,
	}, nil
}
