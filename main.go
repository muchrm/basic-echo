package main

import "github.com/muchrm/go-echo/src/delivery/http"

func main() {
	app := http.New("./config.yml")
	app.Run()
}
