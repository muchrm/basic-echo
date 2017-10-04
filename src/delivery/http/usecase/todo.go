package usecase

import (
	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
)

type Todo interface {
	GetTodoList() (*todoDomain.Todos, error)
}
