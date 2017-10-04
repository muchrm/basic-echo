package usecase

import (
	"testing"
	"time"

	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
	repositoryMock "github.com/muchrm/go-echo/src/usecase/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	t.Parallel()
	repository := &repositoryMock.Todo{}
	todos := &todoDomain.Todos{todoDomain.Todo{"test", false, time.Now()}}
	repository.On(
		"GetTodoList",
	).Return(todos, nil).Times(1)
	interactor := NewTodoInterActor(repository)
	gotTodos, err := interactor.GetTodoList()

	assert.Nil(t, err)
	assert.Equal(t, todos, gotTodos)
	repository.AssertExpectations(t)
}
