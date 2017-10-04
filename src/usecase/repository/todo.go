package repository

import (
	"time"

	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
)

type Todo interface {
	GetTodoList() (*todoDomain.Todos, error)
}
type TodoRepository struct{}

func (r *TodoRepository) List() todoDomain.Todos {
	return todoDomain.Todos{
		todoDomain.Todo{Name: "test", Completed: true, Due: time.Now()},
		todoDomain.Todo{Name: "test", Completed: true, Due: time.Now()},
		todoDomain.Todo{Name: "test", Completed: true, Due: time.Now()},
	}
}
