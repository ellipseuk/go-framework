package main

import (
	"fmt"
	"myframework/internal/core"
)

func main() {
	fmt.Println("Starting the application...")

	app := core.NewApp()

	fmt.Println("Server is running on http://localhost:8080")
	if err := app.Start(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
