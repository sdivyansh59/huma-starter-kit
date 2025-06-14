//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/sdivyansh59/huma-project-starter/app/greeting"
	"github.com/sdivyansh59/huma-project-starter/app/setup"
)

// InitializeApp wires up all dependencies and returns the application/service instance
func InitializeApp() (*App, error) {
	wire.Build(
		setup.ProvideSingletonChiRouter,
		setup.ProvideSingletonHuma,
		setup.ProvideDefaultConfig,
		setup.ProvideControllers,
		newApp,

		// initialize all domains controller and repository
		greeting.NewController,
		//greeting.NewRepository,
	)
	return nil, nil
}
