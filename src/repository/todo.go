package repository

import (
	"time"

	"github.com/muchrm/go-echo/src/domain"
)

type TodoRepository struct{}

func (r *TodoRepository) List() domain.Todos {
	return domain.Todos{
		domain.Todo{Name: "test", Completed: true, Due: time.Now()},
		domain.Todo{Name: "test", Completed: true, Due: time.Now()},
		domain.Todo{Name: "test", Completed: true, Due: time.Now()},
	}
}
