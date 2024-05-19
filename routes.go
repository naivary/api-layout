package main

import (
	"net/http"

	"github.com/naivary/api-layout/handler"
)

func newRouter() http.Handler {
	mux := http.NewServeMux()
	// add your routes here
	mux.Handle("GET /healthz", handler.Healthz())
	mux.Handle("/", http.NotFoundHandler())
	return mux
}
