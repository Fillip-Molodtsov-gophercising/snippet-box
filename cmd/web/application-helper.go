package main

import (
	"fmt"
	"github.com/Fillip-Molodtsov-gophercising/snippet-box/pkg/repository"
	"log"
	"net/http"
	"runtime/debug"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	snippetRepo repository.SnippetGetterUpdater
}

// The serverError helper writes an error message and stack trace to the errorLog,
//then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%v\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to // the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
