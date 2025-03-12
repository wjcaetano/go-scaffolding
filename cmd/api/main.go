package main

import (
	"go-scaffolding/cmd/api/modules"
)

func main() {
	app := modules.NewApp()
	app.Run()
}
