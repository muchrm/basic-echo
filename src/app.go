package src

import "github.com/labstack/echo"

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
	app := &App{
		Engine: engine,
		Config: config,
	}
	return app
}

// Run runs the app
func (app *App) Run() error {
	return app.Engine.Start(":" + app.Config.Port)
}
