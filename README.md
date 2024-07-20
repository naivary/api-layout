# API Layout
This is an api-layout which can be used to implement a RESTful API in Go. The
layout is suggesting tools and methods to be used based on best pratices but is
allowing to expand however the user likes it.

## TODOs
The following list states the requirements which are gradually implemented in
the api-layout.

- [ ] Error handling
- [ ] Route adding
- [ ] Graceful shutdown
- [ ] HTTP/S Certificates
- [ ] Configuration of the API using ENV variables, CLI and config files (12 Factor app)
- [ ] General Logging approach based upon the 12 Factor app method
- [ ] Default Values
- [ ] Static validation via JSON Schema
- [ ] Complex Validation using a common interface
- [ ] Decoding and Encoding

## Error handling
Handling errors in a `HandlerFunc` is challenging. One can't follow the usual
pattern in Go which is to return the error from the function. Having this type
of behavior for the `HandlerFunc` too would be really convenient which is why
the api-layout is introducing the `type HandlerFuncErr func(w http.ResponseWriter,
r *http.Request) error` type. You may know this type from frameworks like echo.
Having this type allows for the usual pattern of Go.

Another challenge is the actual error type. HTTP services return different type
of errors compared to usual Go functions. For example HTTP services have to
return a status code to inform the client how the request went. Fortunately Go
is allowing to easily introduce your own error type. This feature is used to
introduce the `HTTPError` which is a specific error type for HTTP services.

Introcducing both of the types don't change the fact that errors have to be
reported back to client using some kind of http response. Using the
`HandlerFuncErr` type allows to implement a custom error handler for all the
registered endpoints.

## Route adding
Ideally all routes of an API can be overseen in one file. Which the update of
http multiplexer in the standard library it can be used for complex routing. The
api-layout is using the standard library and is allowing to define all the
routes

## Static validation via JSON Schema
