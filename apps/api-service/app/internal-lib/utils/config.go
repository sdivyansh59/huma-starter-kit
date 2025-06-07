package utils

import (
	"fmt"
	"os"
	"strings"
)

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
