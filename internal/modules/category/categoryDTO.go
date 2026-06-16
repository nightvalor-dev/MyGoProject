package category

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

type CategoryResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

func toCategoryResponse(c Category) CategoryResponse {
	return CategoryResponse{
		ID:           c.ID,
		CategoryName: c.CategoryName,
		Description:  c.Description,
	}
}
