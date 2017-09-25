package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", getUser)
	e.Logger.Fatal(e.Start(":3000"))
}
func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
