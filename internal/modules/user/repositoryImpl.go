package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(user User) error {
	query := `INSERT INTO users (username, email, phone, bio, password_hash, role) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := ur.db.Exec(context.Background(), query,
		user.Username, user.Email, user.Phone,
		user.Bio, user.PasswordHash, user.AssignedRole,
	)
	return err
}

func (ur *userRepository) GetAll() ([]User, error) {
	query := `SELECT username, email, phone, bio, role, created_at, updated_at FROM users`
	rows, err := ur.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		var entity User
		err = rows.Scan(
			&entity.Username, &entity.Email,
			&entity.Phone, &entity.Bio, &entity.AssignedRole,
			&entity.CreatedAt, &entity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, entity)
	}
	return result, nil
}

func (ur *userRepository) GetById(id int) (User, error) {
	var entity User
	query := `SELECT username, email, phone, bio, role, created_at, updated_at FROM users WHERE id = $1`
	err := ur.db.QueryRow(context.Background(), query, id).Scan(
		&entity.Username, &entity.Email,
		&entity.Phone, &entity.Bio, &entity.AssignedRole,
		&entity.CreatedAt, &entity.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return entity, nil
}

func (ur *userRepository) Update(id int, newUser User) error {
	args := []any{}
	argIdx := 1
	query := "UPDATE users SET "
	sep := ""

	if newUser.Username != "" {
		query += fmt.Sprintf("%s username = $%d", sep, argIdx)
		args = append(args, newUser.Username)
		argIdx++
		sep = ","
	}

	if newUser.Email != "" {
		query += fmt.Sprintf("%s email = $%d", sep, argIdx)
		args = append(args, newUser.Email)
		argIdx++
		sep = ","
	}

	if newUser.Phone != "" {
		query += fmt.Sprintf("%s phone = $%d", sep, argIdx)
		args = append(args, newUser.Phone)
		argIdx++
		sep = ","
	}

	if newUser.Bio != "" {
		query += fmt.Sprintf("%s bio = $%d", sep, argIdx)
		args = append(args, newUser.Bio)
		argIdx++
		sep = ","
	}

	if len(args) == 0 {
		return nil
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, id)

	_, err := ur.db.Exec(context.Background(), query, args...)
	return err
}

func (ur *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := ur.db.Exec(context.Background(), query, id)
	return err
}

func (ur *userRepository) GetByUsername(username string) (User, error) {
	var entity User
	query := `SELECT id, username, email, phone, bio, password_hash, role, created_at, updated_at FROM users WHERE username = $1`
	err := ur.db.QueryRow(context.Background(), query, username).Scan(
		&entity.ID, &entity.Username, &entity.Email,
		&entity.Phone, &entity.Bio, &entity.PasswordHash,
		&entity.AssignedRole, &entity.CreatedAt, &entity.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return entity, nil
}
