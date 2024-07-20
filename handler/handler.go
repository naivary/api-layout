package handler

import "net/http"

type HandlerFuncErr func(w http.ResponseWriter, r *http.Request) error

type ErrorHandler func(w http.ResponseWriter, err error)

func DefaultErrorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func NewHandler(fn HandlerFuncErr, errFn ErrorHandler) Handler {
    if errFn == nil {
        errFn = DefaultErrorHandler
    }

	return Handler{
        Handler: fn,
		ErrorHandler: errFn,
	}
}

var _ http.Handler = (*Handler)(nil)

type Handler struct {
	Handler HandlerFuncErr

	ErrorHandler func(w http.ResponseWriter, err error)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Handler(w, r)
	if err != nil {
		h.ErrorHandler(w, err)
	}
}
