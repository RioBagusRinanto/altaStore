package main

import (
	"altaStore/config"
	"altaStore/routes"
	"log"
	"os"
)

func main() {

	config.InitDB()

	e := routes.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))

}
