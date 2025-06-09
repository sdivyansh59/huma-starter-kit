package setup

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	appMiddleware "github.com/sdivyansh59/huma-project-starter/app/middleware"
	"sync"
	"time"
)

var (
	routerOnce     sync.Once
	routerInstance *chi.Mux
)

// ProvideSingletonChiRouter returns a singleton chi router
func ProvideSingletonChiRouter() *chi.Mux {
	routerOnce.Do(func() {
		routerInstance = chi.NewRouter()

		// Basic middlewares
		routerInstance.Use(middleware.RequestID)     // Add request ID to each request
		routerInstance.Use(middleware.RealIP)        // Use X-Forwarded-For or X-Real-IP to get client IP
		routerInstance.Use(appMiddleware.ZeroLogger) // Log API request details
		routerInstance.Use(middleware.Recoverer)     // Recover from panics without crashing server

		// Timeout middleware to prevent handlers from running too long
		routerInstance.Use(middleware.Timeout(60 * time.Second))

		// CORS middleware for browser clients
		routerInstance.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
		routerInstance.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
		routerInstance.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization"))

		// Additional useful middlewares
		routerInstance.Use(middleware.CleanPath)    // Clean duplicate slashes in URL paths
		routerInstance.Use(middleware.StripSlashes) // Strip trailing slashes from request paths

		// Optional compression middleware
		routerInstance.Use(middleware.Compress(5)) // Compress responses (level 5)

		routerInstance.Use(appMiddleware.Authenticate)
	})

	return routerInstance
}

// Todo: In your routes setup
//r.Group(func(r chi.Router) {
//	// Apply authentication middleware to this group
//	r.Use(middleware.Authenticate)
//
//	// Protected routes
//	r.Get("/protected-endpoint", protectedHandler)
//})
