package main

import (
	"log"
	"net/http"
	"time"
)

// Um middleware é uma camada de processamento que a mensagem
// vai passar antes de chegar as camadas interiores.
// https://www.alexedwards.net/blog/making-and-using-middleware
func logMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {

		log.Println("message started")
		t := time.Now()

		next.ServeHTTP(w, r)

		log.Println("message completed : latency ->", time.Since(t).String())
	}

	return http.HandlerFunc(middleware)
}
