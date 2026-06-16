package tag

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type tagRepository struct {
	db *pgxpool.Pool
}

func NewTagRepository(db *pgxpool.Pool) TagRepository {
	return &tagRepository{db: db}
}

func (tr *tagRepository) Create(t Tag) error {
	query := `INSERT INTO tags (tag_name) VALUES ($1)`
	_, err := tr.db.Exec(context.Background(), query, t.TagName)
	return err
}

func (tr *tagRepository) GetAll() ([]Tag, error) {
	query := `SELECT tag_name, created_at FROM tags`
	rows, err := tr.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Tag
	for rows.Next() {
		var entity Tag
		err = rows.Scan(&entity.TagName, &entity.Created_at)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}
	return result, nil
}

func (tr *tagRepository) GetById(id int) (Tag, error) {
	var entity Tag
	query := `SELECT tag_name, created_at FROM tags WHERE id = $1`
	err := tr.db.QueryRow(context.Background(), query, id).Scan(&entity.TagName, &entity.Created_at)
	if err != nil {
		return Tag{}, err
	}
	return entity, nil
}

func (tr *tagRepository) Update(id int, newTag Tag) error {
	args := []any{}
	argIdx := 1
	query := "UPDATE tags SET "
	sep := ""

	if newTag.TagName != "" {
		query += fmt.Sprintf("%s tag_name = $%d", sep, argIdx)
		args = append(args, newTag.TagName)
		argIdx++
		sep = ","
	}

	if len(args) == 0 {
		return nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, id)

	_, err := tr.db.Exec(context.Background(), query, args...)
	return err
}

func (tr *tagRepository) Delete(id int) error {
	query := `DELETE FROM tags WHERE id = $1`
	_, err := tr.db.Exec(context.Background(), query, id)
	return err
}

//
//func (tr *tagRepository) GetByPostID(postID int) ([]Tag, error) {
//	query := `SELECT tag_name, blog_id, created_at FROM tags WHERE blog_id = $1`
//	rows, err := tr.db.Query(context.Background(), query, postID)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var result []Tag
//	for rows.Next() {
//		var entity Tag
//		err = rows.Scan(
//			&entity.TagName, &entity.Blog_id, &entity.Created_at,
//		)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, entity)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
