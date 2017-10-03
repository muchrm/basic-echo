package main

import "github.com/muchrm/go-echo/src/app"

func main() {
	app := app.New("./config.yml")
	app.Run()
}
