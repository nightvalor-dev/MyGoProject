package tag

type CreateTagRequest struct {
	TagName string `json:"tag_name"`
	BlogID  int    `json:"blog_id"`
}

type UpdateTagRequest struct {
	TagName string `json:"tag_name"`
}

type TagResponse struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}

func toTagResponse(t Tag) TagResponse {
	return TagResponse{
		ID:      t.ID,
		TagName: t.TagName,
	}
}
