package demodb

import (
	"time"

	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
	"github.com/muchrm/go-echo/src/infrastructure/repository/demodb/transformer"
)

func NewTodo() *Todo {
	return &Todo{
		transformer: transformer.Todo{},
	}
}

// Account provides methods to query user account related data
type Todo struct {
	transformer transformer.Todo
}

func (todo *Todo) GetTodoList() (*todoDomain.Todos, error) {
	return &todoDomain.Todos{todoDomain.Todo{"test", false, time.Now()}}, nil
}
