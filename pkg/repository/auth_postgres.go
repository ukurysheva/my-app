package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	myapp "github.com/ukurysheva/my-app"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "user_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user myapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO  %s (name, username, password_hash) values ($1, $2, $3) RETURNING ID", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (myapp.User, error) {
	var user myapp.User
	query := fmt.Sprintf("SELECT id FROM  %s WHERE username=$1 AND password_hash=$2", usersTable)
	error := r.db.Get(&user, query, username, password)
	return user, error
}
