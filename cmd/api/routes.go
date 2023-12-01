package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	mux.Get("/status", app.status)

	mux.Group(func(mux chi.Router) {
		mux.Use(app.requireBasicAuthentication)

		mux.Get("/basic-auth-protected", app.protected)
	})

	mux.Get("/books", app.getBooks)

	mux.Get("/books/{isbn}", app.getBookByIsbn)

	mux.Get("/books/{id}", app.getBookById)

	// Actions for Librarians
	mux.Post("/books", app.addBook)

	mux.Delete("/books/{id}", app.removeBook)
	return mux
}
