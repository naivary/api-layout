package controller

import "net/http"

type HandlerFuncErr func(w http.ResponseWriter, r *http.Request) error
