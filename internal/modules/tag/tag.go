package tag

import "time"

type Tag struct {
	ID         int       `db:"id"         json:"id"`
	TagName    string    `db:"tag_name"   json:"tag_name"`
	Created_at time.Time `db:"created_at" json:"created_at"`
}
