package blog

type BlogStatus string

const (
	StatusDraft     BlogStatus = "draft"
	StatusPublished BlogStatus = "published"
	StatusArchived  BlogStatus = "archived"
)
