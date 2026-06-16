package comment

import "time"

type Comment struct {
	ID         int       `db:"id" json:"id"`
	Content    string    `db:"content" json:"content"`
	Blog_id    int       `db:"blog_id" json:"blog_id"`
	User_id    int       `db:"user_id" json:"user_id"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
}
