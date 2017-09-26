package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getHome)
	e.Logger.Fatal(e.Start(":3000"))
}

type ResponseContent struct {
	Response  string    `json:"response"`
	Timestamp time.Time `json:"timestamp"`
	Random    int       `json:"random"`
}

func getHome(c echo.Context) error {
	content := ResponseContent{"Sent via JSONP", time.Now().UTC(), rand.Intn(1000)}
	return c.JSON(http.StatusOK, content)
}
