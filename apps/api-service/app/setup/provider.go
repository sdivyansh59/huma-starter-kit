package setup

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/sdivyansh59/huma-project-starter/app/greeting"
	"github.com/sdivyansh59/huma-project-starter/app/internal-lib/utils"
	"sync"
)

var (
	routerOnce     sync.Once
	routerInstance *chi.Mux

	humaOnce     sync.Once
	humaInstance *huma.API
)

// ProvideSingletonChiRouter returns a singleton chi router
func ProvideSingletonChiRouter() *chi.Mux {
	routerOnce.Do(func() {
		routerInstance = chi.NewRouter()
	})
	return routerInstance
}

// ProvideSingletonHuma returns a singleton Huma API instance
func ProvideSingletonHuma(router *chi.Mux) *huma.API {
	humaOnce.Do(func() {
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))
		humaInstance = utils.ToPointer(api)
	})
	return humaInstance
}

// Controllers holds all application controllers
type Controllers struct {
	Greeting *greeting.Controller
	// Add other controllers here as you build them
	// User     *user.Controller
	// Product  *product.Controller
}

// InitializeControllers wires up all controllers
func InitializeControllers(
	greetingController *greeting.Controller,
	// Add other controllers here as parameters
) *Controllers {
	return &Controllers{
		Greeting: greetingController,
		// Add other controllers
	}
}
