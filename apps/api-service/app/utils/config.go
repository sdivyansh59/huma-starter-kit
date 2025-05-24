package utils

import (
	"fmt"
	"os"
	"strings"
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

// GetConfig returns the default config.
func GetConfig(servicePrefix string) *DefaultConfig {
	if servicePrefix != "" {
		servicePrefix += "_"
	}

	port := GetEnvOr(fmt.Sprintf("%sGRPC_PORT", servicePrefix), "8000")
	httpPort := GetEnvOr("PORT", "8001")

	return &DefaultConfig{
		Environment:   GetEnvironment(),
		IsDebug:       IsDebug(),
		APIKey:        MustGetEnv(fmt.Sprintf("%sAPI_KEY", servicePrefix)),
		GrpcAddress:   fmt.Sprintf("localhost:%s", port),
		HTTPAddress:   fmt.Sprintf("0.0.0.0:%s", httpPort),
		Version:       GetEnvOr("VERSION", "1.0"),
		ServicePrefix: servicePrefix,
	}
}

// GetEnvironment returns the current environment.
func GetEnvironment() Environment {
	env := GetEnvOr("ENVIRONMENT", "development")

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
	return GetEnvOr("CI", "false") == trueString
}

// IsDebug returns true if DEBUG is set to true.
func IsDebug() bool {
	return GetEnvOr("DEBUG", trueString) == trueString
}

// GetEnvOr returns the given env variable or a default.
func GetEnvOr(key, otherwise string) string {
	env := os.Getenv(key)

	if env == "" {
		return otherwise
	}

	return env
}

// MustGetEnv returns the given env variable or panics.
func MustGetEnv(key string) string {
	env := os.Getenv(key)

	if env == "" {
		panic(fmt.Sprintf("%s is not set", key))
	}

	return env
}

// GetEnvOrPanicPrefix looks for the env variable prefix_key and returns or panics.
func GetEnvOrPanicPrefix(prefix, key string) string {
	return MustGetEnv(fmt.Sprintf("%s_%s", prefix, key))
}

// GetEnvOrPrefix looks for the env variable prefix_key and returns.
func GetEnvOrPrefix(prefix, key, or string) string {
	if prefix == "" {
		return GetEnvOr(key, or)
	}

	return GetEnvOr(fmt.Sprintf("%s_%s", prefix, key), or)
}

// SplitSimpleConfig splits a simple config string into a map.
// The config is of the form `key1=value1,key2=value2`.
func SplitSimpleConfig(input string) map[string]string {
	config := make(map[string]string)

	if input == "" {
		return config
	}

	split := strings.Split(input, ",")

	for _, s := range split {
		split2 := strings.Split(s, "=")

		if len(split2) != 2 {
			continue
		}

		config[split2[0]] = split2[1]
	}

	return config
}
