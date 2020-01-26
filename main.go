package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// main
func main() {
	log.Info("This is a log line")
	log.Warn("Another log line")
	log.Error("This is a really bad")

	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "package main #14")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("url: %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
