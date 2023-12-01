package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nataliabu/library-service/internal/database"
	"github.com/nataliabu/library-service/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {
	responseBody, err := listBooksDB(app.db, r.Context())
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getBookByIsbn(w http.ResponseWriter, r *http.Request) {
	isbn := chi.URLParam(r, "isbn")
	responseBody, err := getBookByIsbnDB(app.db, r.Context(), &isbn)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getBookById(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	parsedId, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	id := int32(parsedId)
	responseBody, err := getBookByIdDB(app.db, r.Context(), &id)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, responseBody)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var book database.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	responseBody, err := addBookDB(app.db, r.Context(), &book)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusCreated, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) removeBook(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	parsedId, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	id := int32(parsedId)
	err = removeBookDB(app.db, r.Context(), &id)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, "Book succesfully deleted")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getCustomers(w http.ResponseWriter, r *http.Request) {
	responseBody, err := listCustomersDB(app.db, r.Context())
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) addCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var customer database.Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	responseBody, err := addCustomerDB(app.db, r.Context(), &customer)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusCreated, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) borrowBook(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	parsedId, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	id := int32(parsedId)

	book, err := getBookByIdDB(app.db, r.Context(), &id)
	if err != nil {
		app.notFound(w, r)
		return
	}

	if !book.Available {
		err = response.JSON(w, http.StatusNotFound, "This book is not available at the moment")
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	} else {
		err := borrowBookDB(app.db, r.Context(), book)
		if err != nil {
			app.notFound(w, r)
			return
		}

		err = response.JSON(w, http.StatusOK, "You've borrowed this book!")
		if err != nil {
			app.serverError(w, r, err)
		}
	}
}

func (app *application) returnBook(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	parsedId, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	id := int32(parsedId)

	book, err := getBookByIdDB(app.db, r.Context(), &id)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = returnBookDB(app.db, r.Context(), book)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = response.JSON(w, http.StatusOK, "You've returned this book. Thank you!")
	if err != nil {
		app.serverError(w, r, err)
	}
}
