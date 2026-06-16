package blog

import (
	"time"
)

type Blog struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Content     string     `db:"content" json:"content"`
	UserId      int        `db:"user_id" json:"user_id"`
	Status      BlogStatus `db:"status" json:"status"`
	CategoryId  int        `db:"category_id" json:"category_id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	PublishedAt time.Time  `db:"published_at" json:"published_at"`
}
