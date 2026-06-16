package blogtag

type BlogTag struct {
	BlogID int `json:"blog_id" db:"blog_id"`
	TagID  int `json:"tag_id"  db:"tag_id"`
}
