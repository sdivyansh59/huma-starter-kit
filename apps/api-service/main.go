package main

import (
	"github.com/getsentry/sentry-go"
	"time"
	_ "time/tzdata" // ensure we always have the timezone information included
)

func main() {
	defer sentry.Flush(2 * time.Second)

}
