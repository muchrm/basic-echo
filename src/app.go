package src

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/muchrm/go-echo/src/handle"
)

// App struct.
type App struct {
	Engine *echo.Echo
	Config *Config
}

// Init struct
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
	handle.NewTodoHandle(
		app.Engine.Group(
			app.Config.Api.Prefix,
		),
	)
	handle.NewStatsHandle(
		app.Engine.Group(
			app.Config.Api.Prefix,
		),
	)
	return app
}

// Run runs the app
func (app *App) Run() error {
	return app.Engine.Start(":" + app.Config.Port)
}
