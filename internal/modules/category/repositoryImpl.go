package category

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type catRepository struct {
	db *pgxpool.Pool
}

func NewCatRepository(db *pgxpool.Pool) CategoryRepository {
	return &catRepository{db: db}
}

func (cat *catRepository) Create(category Category) error {
	query := `INSERT INTO categories (category_name, description) VALUES ($1, $2)`
	_, err := cat.db.Exec(context.Background(), query, category.CategoryName, category.Description)
	return err
}

func (cat *catRepository) GetAll() ([]Category, error) {
	query := "SELECT id, category_name, description, created_at FROM categories"
	rows, err := cat.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Category
	for rows.Next() {
		var entity Category
		err = rows.Scan(&entity.ID, &entity.CategoryName, &entity.Description, &entity.Created_at)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}

	return result, nil
}

func (cat *catRepository) GetById(id int) (Category, error) {
	var entity Category
	query := `SELECT id, category_name, description, created_at FROM categories WHERE id = $1`
	err := cat.db.QueryRow(context.Background(), query, id).Scan(&entity.ID,
		&entity.CategoryName, &entity.Description, &entity.Created_at)

	if err != nil {
		return Category{}, err
	}

	return entity, nil
}

func (cat *catRepository) Update(id int, newCategory Category) error {
	args := []any{}
	argIdx := 1
	query := "UPDATE categories SET "
	sep := ""

	if newCategory.CategoryName != "" {
		query += fmt.Sprintf("%s category_name = $%d", sep, argIdx)
		args = append(args, newCategory.CategoryName)
		argIdx++
		sep = ","
	}

	if newCategory.Description != "" {
		query += fmt.Sprintf("%s description = $%d", sep, argIdx)
		args = append(args, newCategory.Description)
		argIdx++
		sep = ","
	}

	if len(args) == 0 {
		return nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, id)

	_, err := cat.db.Exec(context.Background(), query, args...)
	return err
}

func (cat *catRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := cat.db.Exec(context.Background(), query, id)
	return err
}

//func (cat *catRepository) GetByPostId(blogID int) (*Category, error) {
//	var entity Category
//	query := `SELECT category_name, description, created_at FROM categories WHERE blog_id = $1`
//	err := cat.db.QueryRow(context.Background(), query, blogID).Scan(
//		&entity.ID, &entity.CategoryName, &entity.Description, &entity.Created_at,
//	)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &entity, nil
//}
