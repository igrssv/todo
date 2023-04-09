package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}
