package main

import (
	"fmt"
	"myframework/internal/core"
	"myframework/internal/modules/example_module"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")

	app := core.NewApp()

	app.Router.Handle("/example", http.HandlerFunc(example_module.ExampleModule))

	if err := app.Start(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
