package handle

import (
	"net/http"

	"github.com/labstack/echo"
	todoUsecase "github.com/muchrm/go-echo/src/delivery/http/usecase"
)

type TodoHandle struct {
	todoUsecase todoUsecase.Todo
	group       *echo.Group
}
type TodoError struct{}

func (handle *TodoHandle) bind() {
	handle.group.GET("/todos", handle.getTodo)
}

func (handle *TodoHandle) getTodo(c echo.Context) error {
	todos, err := handle.todoUsecase.GetTodoList()
	if err != nil {
		return c.JSON(http.StatusNotFound, TodoError{})
	} else {
		return c.JSON(http.StatusOK, todos)
	}
}

func NewTodoHandle(todoUsecase todoUsecase.Todo, group *echo.Group) *TodoHandle {
	handle := TodoHandle{todoUsecase, group}
	handle.bind()
	return &handle
}
