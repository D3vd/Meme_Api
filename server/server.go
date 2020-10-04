package server

import "github.com/getsentry/sentry-go"

// Init : Initialize the routes and server
func Init() {
	r := NewRouter()
	err := r.Run()

	if err != nil {
		sentry.CaptureException(err)
	}
}
