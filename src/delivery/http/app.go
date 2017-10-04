package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/muchrm/go-echo/src/delivery/http/handle"
	"github.com/muchrm/go-echo/src/infrastructure/repository/demodb"
	todoUsecase "github.com/muchrm/go-echo/src/usecase"
)

// App struct.
type App struct {
	Engine *echo.Echo
	Config *Config
}

// New struct
func New(configPath string) *App {
	engine := echo.New()
	config, err := ParseYaml(configPath)
	// Initialize the application
	if err != nil {
		panic(err)
	}
	engine.Use(middleware.Recover())
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
	}))

	app := &App{
		Engine: engine,
		Config: config,
	}
	todoRepository := demodb.NewTodo()
	todoInteractor := todoUsecase.NewTodoInterActor(todoRepository)
	handle.NewTodoHandle(todoInteractor, app.Engine.Group(app.Config.Api.Prefix))
	return app
}

// Run runs the app
func (app *App) Run() error {
	return app.Engine.Start(":" + app.Config.Port)
}
