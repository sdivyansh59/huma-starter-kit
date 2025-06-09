package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/sdivyansh59/huma-project-starter/app"
	appMiddleware "github.com/sdivyansh59/huma-project-starter/app/middleware"
	"time"
	_ "time/tzdata" // ensure we always have the timezone information included
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	appMiddleware.InitZeroLog()
}

func main() {
	defer sentry.Flush(2 * time.Second)

	application, err := app.InitializeApp()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	if err := application.Run(); err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}
