package handle

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muchrm/go-echo/src/repository"
)

func NewTodoHandle(group *echo.Group) {
	group.GET("/todos", getTodo)
}

func getTodo(c echo.Context) error {
	r := repository.TodoRepository{}
	return c.JSON(http.StatusOK, r.List())
}
