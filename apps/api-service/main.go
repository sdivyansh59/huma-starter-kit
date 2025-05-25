package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/sdivyansh59/huma-project-starter/app"
	"github.com/sdivyansh59/huma-project-starter/app/setup"
	"net/http"
	"time"
	_ "time/tzdata" // ensure we always have the timezone information included

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func main() {
	defer sentry.Flush(2 * time.Second)

	// Create a new router & API
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	app.RegisterRoutes(api)
	_, err := setup.InitializeDatabase()
	if err != nil {
		sentry.CaptureException(err)
		return
	}

	// Start the server!
	err = http.ListenAndServe("127.0.0.1:8888", router)
	if err != nil {
		return
	}
}
