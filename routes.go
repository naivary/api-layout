package main

import (
	"net/http"

	"github.com/naivary/api-layout/controller"
)

func newRouter() http.Handler {
	mux := http.NewServeMux()
    // TODO(user): Add your routes here. `mux` is the standard multiplexer
    // from the standard library and can be exchanged with any type of
    // multiplexer you which to use.

	// mux.Handle("GET /healthz", handler.Healthz())
    mux.Handle("GET /healthz", controller.HandleSomething())
	return mux
}
