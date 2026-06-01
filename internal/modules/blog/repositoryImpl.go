package blog

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type blogRepository struct {
	db *pgxpool.Pool
}

func NewBlogRepository(db *pgxpool.Pool) BlogRepository {
	return &blogRepository{db: db}
}

func (b *blogRepository) Create(blog Blog) (int, error) {
	var id int
	query := `INSERT INTO blogs (title, content, user_id, status) VALUES ($1, $2, $3, $4) RETURNING id`
	err := b.db.QueryRow(context.Background(), query, blog.Title, blog.Content, blog.User_id, blog.Status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *blogRepository) GetAll() ([]Blog, error) {
	query := `SELECT title, content, user_id, status, created_at, updated_at, published_at FROM blogs`
	rows, err := b.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Blog
	for rows.Next() {
		var entity Blog
		err = rows.Scan(
			&entity.Title, &entity.Content,
			&entity.User_id, &entity.Status,
			&entity.Created_at, &entity.Updated_at, &entity.Published_at,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}
	return result, nil
}

func (b *blogRepository) GetById(id int) (Blog, error) {
	var entity Blog
	query := `SELECT title, content, user_id, status, created_at, updated_at, published_at FROM blogs WHERE id = $1`
	err := b.db.QueryRow(context.Background(), query, id).Scan(
		&entity.Title, &entity.Content,
		&entity.User_id, &entity.Status,
		&entity.Created_at, &entity.Updated_at, &entity.Published_at,
	)
	if err != nil {
		return Blog{}, err
	}
	return entity, nil
}

func (b *blogRepository) Update(id int, newBlog Blog) error {
	query := `
		UPDATE blogs
		SET
			title   = CASE WHEN $1::text IS NOT NULL THEN $1::text ELSE title END,
			content = CASE WHEN $2::text IS NOT NULL THEN $2::text ELSE content END,
			status  = CASE WHEN $3::text IS NOT NULL THEN $3::text ELSE status END
		WHERE id = $4`
	_, err := b.db.Exec(context.Background(), query, newBlog.Title, newBlog.Content, newBlog.Status, id)
	return err
}

func (b *blogRepository) Delete(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`
	_, err := b.db.Exec(context.Background(), query, id)
	return err
}
