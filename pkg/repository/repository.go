package repository

import (
	"github.com/jmoiron/sqlx"
	myapp "github.com/ukurysheva/my-app"
)

type Authorization interface {
	CreateUser(user myapp.User) (int, error)
	GetUser(username, password string) (myapp.User, error)
}

type TodoList interface {
}
type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
