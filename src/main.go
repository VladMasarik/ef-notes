package main

import (
	"fmt"

	"github.com/VladMasarik/ef-notes/src/database"
	"github.com/VladMasarik/ef-notes/src/routes"
)

func main() {
	if err := database.SetupDatabase(); err != nil {
		panic("Failed to setup a database.")
	}
	r := routes.SetupRoutes()

	fmt.Println("Setup succeed, running server...")
	r.Run()
}
