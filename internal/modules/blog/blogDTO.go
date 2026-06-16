package blog

// For POST operation
type CreateBlogRequest struct {
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	UserId   int               `json:"user_id"`
	Status   string            `json:"status"`
	Category CreateCategoryDTO `json:"category"`
	Tags     []CreateTagDTO    `json:"tags"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
type CreateCategoryDTO struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
type CreateTagDTO struct {
	TagName string `json:"tag_name"`
}

// For GET operation
type BlogResponse struct {
	ID       int           `json:"id"`
	Title    string        `json:"title"`
	Content  string        `json:"content"`
	UserId   int           `json:"user_id"`
	Status   string        `json:"status"`
	Category CategoryLocal `json:"category"`
	Tags     []TagLocal    `json:"tags"`
}

// Locally stored in the blog module so there is no relation b/w any modules
type CategoryLocal struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
type TagLocal struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}
