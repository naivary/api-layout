package main

import (
	"net"
	"net/http"
)

func NewHTTPServer() *http.Server {
    mux := newRouter()
    server := &http.Server{
        Addr: net.JoinHostPort("localhost", "8080"),
        Handler: mux,
    }
    return server
}
