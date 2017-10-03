package handle

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muchrm/go-echo/src/repository"
)

type TodoHandle struct {
	todoRepository *repository.TodoRepository
	group          *echo.Group
}
func (handle *TodoHandle) bind() {
	handle.group.GET("/todos", handle.getTodo)
}

func (handle *TodoHandle) getTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, handle.todoRepository.List())
}

func NewTodoHandle(todoRepository *repository.TodoRepository,group *echo.Group) *TodoHandle{
	handle := TodoHandle{todoRepository,group}
	handle.bind();
	return &handle
}

