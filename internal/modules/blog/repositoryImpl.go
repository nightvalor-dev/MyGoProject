package blog

import (
	"context"
	"fmt"

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
	query := `INSERT INTO blogs (title, content, user_id, category_id, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := b.db.QueryRow(context.Background(), query, blog.Title, blog.Content,
		blog.UserId, blog.CategoryId, blog.Status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *blogRepository) GetAll() ([]Blog, error) {
	query := "SELECT title, content, status FROM blogs"
	rows, err := b.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Blog
	for rows.Next() {
		var entity Blog
		err = rows.Scan(&entity.Title, &entity.Content)
		if err != nil {
			return nil, err
		}

		result = append(result, entity)
	}
	return result, nil
}

func (b *blogRepository) GetById(id int) (Blog, error) {
	var entity Blog
	query := `SELECT title, content, status FROM blogs WHERE id = $1`
	err := b.db.QueryRow(context.Background(), query, id).Scan(&entity.Title, &entity.Content)
	if err != nil {
		return Blog{}, err
	}
	return entity, nil
}

func (b *blogRepository) Update(id int, newBlog Blog) error {
	args := []any{}
	argIdx := 1
	query := "UPDATE blogs SET "
	sep := ""

	if newBlog.Title != "" {
		query += fmt.Sprintf("%s title = $%d", sep, argIdx)
		args = append(args, newBlog.Title)
		argIdx++
		sep = ","
	}

	if newBlog.Content != "" {
		query += fmt.Sprintf("%s content = $%d", sep, argIdx)
		args = append(args, newBlog.Content)
		argIdx++
		sep = ","
	}

	if newBlog.Status != "" {
		query += fmt.Sprintf("%s status = $%d", sep, argIdx)
		args = append(args, newBlog.Status)
		argIdx++
		sep = ","
	}

	if len(args) == 0 {
		return nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, id)

	_, err := b.db.Exec(context.Background(), query, args...)
	return err
}

func (b *blogRepository) Delete(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`
	_, err := b.db.Exec(context.Background(), query, id)
	return err
}
