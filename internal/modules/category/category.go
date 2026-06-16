package category

import "time"

type Category struct {
	ID           int       `db:"id"            json:"id"`
	CategoryName string    `db:"category_name" json:"category_name"`
	Description  string    `db:"description"   json:"description"`
	Created_at   time.Time `db:"created_at"    json:"created_at"`
}
