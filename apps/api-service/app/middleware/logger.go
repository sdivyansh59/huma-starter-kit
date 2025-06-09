package middleware

import (
	"github.com/sdivyansh59/huma-project-starter/app/internal-lib/utils"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitZeroLog configures zerolog settings
func InitZeroLog() {
	// Determine environment
	env := utils.GetEnvOr("ENVIRONMENT", "production")
	isProd := env == "production"

	// Configure log level
	logLevel := zerolog.InfoLevel
	if debugMode := os.Getenv("DEBUG"); debugMode == "true" {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	// Get hostname, with fallback
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// Configure output writer
	var output io.Writer = os.Stdout
	if isProd {
		// JSON format for production (better for log aggregators)
		output = os.Stdout
	} else {
		// Pretty console output for development
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006/01/02 15:04:05",
		}
	}

	// Set up the logger
	log.Logger = zerolog.New(output).
		With().
		Timestamp().
		Str("host", hostname).
		Str("env", env).
		Logger()

	log.Info().Msg("Logger initialized")
}

// ZeroLogger is a middleware that logs HTTP requests using zerolog
func ZeroLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			duration := time.Since(start)
			requestID := middleware.GetReqID(r.Context())

			// Fully structured approach
			log.Info().
				Str("requestID", requestID).
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("proto", r.Proto).
				Str("remoteAddr", r.RemoteAddr).
				Int("status", ww.Status()).
				Int("size", ww.BytesWritten()).
				Dur("duration", duration).
				Msg("request completed")
		}()

		next.ServeHTTP(ww, r)
	})
}
