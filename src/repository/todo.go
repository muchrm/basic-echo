package repository

import (
	"time"

	"github.com/muchrm/go-echo/src/domain/todo"
)

type TodoRepository struct{}

func (r *TodoRepository) List() todo.Todos {
	return todo.Todos{
		todo.Todo{Name: "test", Completed: true, Due: time.Now()},
		todo.Todo{Name: "test", Completed: true, Due: time.Now()},
		todo.Todo{Name: "test", Completed: true, Due: time.Now()},
	}
}
