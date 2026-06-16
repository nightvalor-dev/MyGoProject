package comment

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type comRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) CommentRepository {
	return &comRepository{db: db}
}

func (com *comRepository) Create(comment Comment) error {
	query := `INSERT INTO comments (content, blog_id, user_id) VALUES ($1, $2, $3)`
	_, err := com.db.Exec(context.Background(), query, comment.Content, comment.Blog_id, comment.User_id)
	return err
}

func (com *comRepository) GetAll() ([]Comment, error) {
	query := `SELECT content, blog_id, user_id, created_at, updated_at FROM comments`
	rows, err := com.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Comment
	for rows.Next() {
		var entity Comment
		err = rows.Scan(
			&entity.Content, &entity.Blog_id, &entity.User_id,
			&entity.Created_at, &entity.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}
	return result, nil
}

func (com *comRepository) GetById(id int) (Comment, error) {
	var entity Comment
	query := `SELECT content, blog_id, user_id, created_at, updated_at FROM comments WHERE id = $1`
	err := com.db.QueryRow(context.Background(), query, id).Scan(
		&entity.Content, &entity.Blog_id, &entity.User_id,
		&entity.Created_at, &entity.Updated_at,
	)
	if err != nil {
		return Comment{}, err
	}
	return entity, nil
}

func (com *comRepository) Update(id int, newComment Comment) error {
	args := []any{}
	argIdx := 1
	query := "UPDATE comments SET "
	sep := ""

	if newComment.Content != "" {
		query += fmt.Sprintf("%s content = $%d", sep, argIdx)
		args = append(args, newComment.Content)
		argIdx++
		sep = ","
	}

	if len(args) == 0 {
		return nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, id)

	_, err := com.db.Exec(context.Background(), query, args...)
	return err
}

func (com *comRepository) Delete(id int) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := com.db.Exec(context.Background(), query, id)
	return err
}
