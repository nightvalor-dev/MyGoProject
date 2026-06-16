package comment

type CreateCommentRequest struct {
	Content string `json:"content"`
	BlogID  int    `json:"blog_id"`
	UserID  int    `json:"user_id"`
}

type UpdateCommentRequest struct {
	Content string `json:"content"`
}

type CommentResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	BlogID  int    `json:"blog_id"`
	UserID  int    `json:"user_id"`
}

func toCommentResponse(c Comment) CommentResponse {
	return CommentResponse{
		ID:      c.ID,
		Content: c.Content,
		BlogID:  c.Blog_id,
		UserID:  c.User_id,
	}
}
