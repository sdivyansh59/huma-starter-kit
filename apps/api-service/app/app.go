package app

import (
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/sdivyansh59/huma-project-starter/app/setup"
	"github.com/sdivyansh59/huma-project-starter/routes"
	"github.com/uptrace/bun"
	"log"
	"net/http"
)

// App is the main application struct
type App struct {
	router      *chi.Mux
	huma        *huma.API
	db          *bun.DB
	controllers *setup.Controllers
	config      *setup.DefaultConfig
}

func newApp(r *chi.Mux, h *huma.API, config *setup.DefaultConfig, c *setup.Controllers) *App {
	db, err := setup.InitializeDatabase()
	if err != nil || db == nil {
		panic(fmt.Sprintf("failed to initialize database: %v", err))
	}

	return &App{
		router:      r,
		huma:        h,
		db:          db,
		controllers: c,
		config:      config,
	}
}

// Run starts the application server
func (a *App) Run() error {
	// Configure your routes
	a.registerRoutes()

	// Start the HTTP server
	log.Printf("Starting server on %s", a.config.HTTPAddress)
	return http.ListenAndServe(a.config.HTTPAddress, a.router)
}

// registerRoutes configures all API endpoints
func (a *App) registerRoutes() {
	if a.huma == nil {
		log.Fatal("huma not initialized")
	}

	routes.RegisterRoutes(a.huma, a.controllers)
}
