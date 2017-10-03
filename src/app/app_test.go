package app

import (
	"reflect"
	"testing"

	"github.com/muchrm/go-echo/src/config"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
)

func TestNew(t *testing.T) {
	confTest := &config.Config{Port: "3000", Api: config.ApiConfig{Prefix: "/api"}}
	appTest := &App{Engine: echo.New(), Config: confTest}
	app := New("../../config.yml")
	if reflect.DeepEqual(appTest, confTest) {
		t.Errorf("TestInit %v got %v", appTest, app)
	}
}
func TestRun(t *testing.T) {
	app := New("../../config.yml")
	go func() {
		assert.NoError(t, app.Run())
	}()
}
