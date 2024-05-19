package handler

import "net/http"

func Healthz() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // your health logic here (highly recommended)
    })
}
