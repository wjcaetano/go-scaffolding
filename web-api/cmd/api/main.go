package main

import (
	"go-scaffolding/web-api/cmd/api/modules"
)

func main() {
	app := modules.NewApp()
	app.Run()
}
