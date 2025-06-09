package setup

import (
	"fmt"
	"github.com/sdivyansh59/huma-project-starter/app/internal-lib/utils"
)

// DefaultConfig defines configuration data that is often used by all services.
type DefaultConfig struct {
	// Environment the service runs on currently.
	Environment Environment

	// IsDebug indicates whether the service is running in debug mode.
	IsDebug bool

	// APIKey is the apikey the service uses for its gRPC server.
	APIKey string

	// GrpcAddress is the address where the gRPC server should listen on.
	GrpcAddress string

	// HttpAddress is the address where the "normal" HTTP server should listen on.
	HTTPAddress string

	// Version is an arbitrary version string for this instance.
	Version string

	// ServiceName defines how the service is called.
	ServiceName string

	// ServicePrefix defines which prefix is used for getting environment variables.
	// For example "SERVICE" would look for "SERVICE_PORT".
	ServicePrefix string
}

type Environment string

const (
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

const trueString = "true"

// ProvideDefaultConfig returns the default config.
func ProvideDefaultConfig() *DefaultConfig {
	servicePrefix := utils.GetEnvOr("SERVICE_PREFIX", "")
	if servicePrefix != "" {
		servicePrefix += "_"
	}

	port := utils.GetEnvOr(fmt.Sprintf("%sGRPC_PORT", servicePrefix), "8000")
	httpPort := utils.GetEnvOr("PORT", "8001")
	apiKey := utils.GetEnvOr("API_KEY", "")

	return &DefaultConfig{
		Environment:   GetEnvironment(),
		IsDebug:       IsDebug(),
		APIKey:        apiKey, //utils.MustGetEnv(fmt.Sprintf("%sAPI_KEY", servicePrefix)),
		GrpcAddress:   fmt.Sprintf("localhost:%s", port),
		HTTPAddress:   fmt.Sprintf("0.0.0.0:%s", httpPort),
		Version:       utils.GetEnvOr("VERSION", "1.0"),
		ServicePrefix: servicePrefix,
	}
}

// GetEnvironment returns the current environment.
func GetEnvironment() Environment {
	env := utils.GetEnvOr("ENVIRONMENT", "development")

	switch env {
	case "development":
		return Development
	case "staging":
		return Staging
	case "production":
		return Production
	default:
		panic("unknown environment")
	}
}

// IsCI returns true if we are running in a CI environment.
func IsCI() bool {
	return utils.GetEnvOr("CI", "false") == trueString
}

// IsDebug returns true if DEBUG is set to true.
func IsDebug() bool {
	return utils.GetEnvOr("DEBUG", trueString) == trueString
}
