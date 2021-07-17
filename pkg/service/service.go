package service

import (
	myapp "github.com/ukurysheva/my-app"
	"github.com/ukurysheva/my-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user myapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
