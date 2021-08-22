package main

import (
	"altaStore/config"
	"altaStore/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
