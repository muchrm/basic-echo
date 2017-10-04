package usecase

import (
	"fmt"

	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
	"github.com/muchrm/go-echo/src/usecase/repository"
)

func NewTodoInterActor(todoRepository repository.Todo) *TodoInterActor {
	return &TodoInterActor{
		repository: todoRepository,
	}
}

type TodoInterActor struct {
	repository repository.Todo
}

func (interactor *TodoInterActor) GetTodoList() (*todoDomain.Todos, error) {
	todos, err := interactor.repository.GetTodoList()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch todo data: %s", err)
	}
	return todos, nil
}
