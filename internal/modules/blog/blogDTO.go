package blog

type CreateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	Status  string `json:"status"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type BlogDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	Status  string `json:"status"`
}
