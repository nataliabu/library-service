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

	mux.Get("/books", app.getBooks)

	mux.Get("/books/{isbn}", app.getBookByIsbn)

	mux.Get("/books/{id}", app.getBookById)

	// Actions for Librarians
	mux.Group(func(mux chi.Router) {
		mux.Use(app.adminBasicAuthentication)

		mux.Post("/books", app.addBook)

		mux.Delete("/books/{id}", app.removeBook)

		mux.Get("/customers", app.getCustomers)

		mux.Post("/customers", app.addCustomer)

	})

	// Actions for Customers
	mux.Group(func(mux chi.Router) {
		mux.Use(app.customerBasicAuthentication)

		mux.Patch("/books/borrow/{id}", app.borrowBook)

		mux.Patch("/books/return/{id}", app.returnBook)
	})

	return mux
}
