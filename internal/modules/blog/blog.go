package blog

import "time"

type Blog struct {
	ID           int        `db:"id" json:"id"`
	Title        string     `db:"title" json:"title"`
	Content      string     `db:"content" json:"content"`
	User_id      int        `db:"user_id" json:"user_id"`
	Status       BlogStatus `db:"status" json:"status"`
	Created_at   time.Time  `db:"created_at" json:"created_at"`
	Updated_at   time.Time  `db:"updated_at" json:"updated_at"`
	Published_at time.Time  `db:"published_at" json:"published_at"`
}
