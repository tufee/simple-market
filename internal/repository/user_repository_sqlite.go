package repository

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
	"simple.market/internal/domain"
)

type userRepositorySQLite struct {
	db *sql.DB
}

func NewUserRepositorySQLite(db *sql.DB) UserRepository {
	return &userRepositorySQLite{db: db}
}

func (r *userRepositorySQLite) Create(user *domain.User) error {
	query := `INSERT INTO users (id, email, password) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, user.ID, user.Email, user.Password)
	return err
}

func (r *userRepositorySQLite) FindByID(id int) (*domain.User, error) {
	query := `SELECT id, email FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepositorySQLite) Update(user *domain.User) error {
	query := `UPDATE users SET email = ? WHERE id = ?`
	result, err := r.db.Exec(query, user.Email, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepositorySQLite) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
