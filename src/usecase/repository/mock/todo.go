package mock

import (
	todoDomain "github.com/muchrm/go-echo/src/domain/todo"
	"github.com/stretchr/testify/mock"
)

type Todo struct {
	mock.Mock
}

func (m *Todo) GetTodoList() (*todoDomain.Todos, error) {
	args := m.Called()
	return args.Get(0).(*todoDomain.Todos), args.Error(1)
}
