package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-web/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Create(payload *model.User) error {
	query := `INSERT INTO users(username, password, email) VALUES($1, $2, $3)`
	_, err := r.db.Exec(query, payload.Username, payload.Password, payload.Email)
	if err != nil {
		return err
	}
	return nil
}

func (cr *UserRepository) UserByID(id int) (*model.User, error) {
	user := model.User{}
	query := `SELECT username, password, email FROM users WHERE id=$1`
	err := cr.db.QueryRow(query, id).Scan(&user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUser() ([]*model.User, error) {
	query := `SELECT id, username, password, email FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return users, nil
}

func (r *UserRepository) Login(customer model.User) error {
	user := model.User{}
	query := `SELECT id, username, password, email FROM users WHERE username=$1 AND password=$2`
	err := r.db.QueryRow(query, customer.Username, customer.Password).Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return fmt.Errorf("failed to execute query %v", err)
	}
	fmt.Println(err)
	return nil
}
