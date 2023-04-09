package servise

import "todo/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
}

func NewServive(repository *repository.Repository) *Service {
	return &Service{}
}
