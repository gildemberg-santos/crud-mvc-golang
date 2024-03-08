package controller

import (
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		log.Printf("%s %s %s %v \n", r.RemoteAddr, r.Method, r.URL, duration)
	})
}
