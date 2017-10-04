package handle

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/muchrm/go-echo/src/infrastructure/repository/demodb"
	todoUsecase "github.com/muchrm/go-echo/src/usecase"
)

func TestNewTodoHandle(t *testing.T) {
	todoRepository := demodb.NewTodo()
	todoInteractor := todoUsecase.NewTodoInterActor(todoRepository)
	NewTodoHandle(todoInteractor, echo.New().Group("/api"))
	t.Skip("skill")
}
