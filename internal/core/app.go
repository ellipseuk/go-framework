package core

import (
	"net/http"
)

// App is the main application struct
type App struct {
	Router *Router
}

// NewApp creates a new App
func NewApp() *App {
	return &App{
		Router: NewRouter(),
	}
}

// Start starts the application
func (a *App) Start(addr string) error {
	return http.ListenAndServe(addr, a.Router)
}
