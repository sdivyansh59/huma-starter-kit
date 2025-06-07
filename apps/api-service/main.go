package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/sdivyansh59/huma-project-starter/app"
	"time"
	_ "time/tzdata" // ensure we always have the timezone information included
)

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
