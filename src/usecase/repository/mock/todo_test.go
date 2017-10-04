package mock

import (
	"testing"
	"time"

	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
)

func TestGetTodoList(t *testing.T) {
	testObj := new(Todo)

	// setup expectations
	todos := &todoDomain.Todos{todoDomain.Todo{"test", false, time.Now()}}
	testObj.On(
		"GetTodoList",
	).Return(todos, nil).Times(1)
	testObj.GetTodoList()
	// assert that the expectations were met
	testObj.AssertExpectations(t)
}
