package blogtag

import (
	"Project2-v3/internal/modules/tag"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type blogTagRepository struct {
	db *pgxpool.Pool
}

func (r *blogTagRepository) AssignTagToBlog(blogID, tagID int) error {
	query := "INSERT INTO blog_tags (blog_id, tag_id) VALUES ($1, $2)"
	_, err := r.db.Exec(context.Background(), query, blogID, tagID)
	return err
}

func (r *blogTagRepository) RemoveTagFromBlog(blogID, tagID int) error {
	query := "DELETE FROM blog_tags WHERE blog_id = $1 AND tag_id = $2"
	_, err := r.db.Exec(context.Background(), query, blogID, tagID)
	return err
}

func (r *blogTagRepository) GetTagsByBlogID(blogID int) ([]tag.Tag, error) {
	query := `
		SELECT t.id, t.tag_name, t.created_at FROM tags t
		INNER JOIN blog_tags bt ON t.id = bt.tag_id
		WHERE bt.blog_id = $1
	`

	rows, err := r.db.Query(context.Background(), query, blogID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tags []tag.Tag
	for rows.Next() {
		var tag tag.Tag

		err = rows.Scan(&tag.ID, &tag.TagName, &tag.Created_at)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, rows.Err()
}
