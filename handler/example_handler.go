package handler

import "net/http"

func HandleSomething() Handler {
	fn := HandlerFuncErr(func(w http.ResponseWriter, r *http.Request) error {
        //TODO(user): your handler logic here
        return nil
    })

	return NewHandler(fn, nil)
}
