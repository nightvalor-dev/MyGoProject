package blogtag

import (
	"Project2-v3/internal/modules/blog"
	"Project2-v3/internal/modules/tag"
)

type BlogTagRepository interface {
	AssignTagToBlog(blogID, tagID int) error
	RemoveTagFromBlog(blogID, tagID int) error
	GetTagsByBlogID(blogID int) ([]tag.Tag, error)
	GetBlogsByTagID(tagID int) ([]blog.Blog, error)
}
