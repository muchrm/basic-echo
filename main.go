package main

import "github.com/muchrm/go-echo/src"

func main() {
	app := src.New("./config.yml")
	app.Run()
}
