package main

import "net/http"

func exampleMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // add your middleware logic here and call
        // the next using `h.ServerHTTP`
        h.ServeHTTP(w, r)
    })
}
