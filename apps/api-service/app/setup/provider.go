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
	humaOnce     sync.Once
	humaInstance *huma.API
)

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

// ProvideControllers wires up all controllers
func ProvideControllers(
	greetingController *greeting.Controller,
// Add other controllers here as parameters
) *Controllers {
	return &Controllers{
		Greeting: greetingController,
		// Add other controllers
	}
}
