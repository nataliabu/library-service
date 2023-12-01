# BOOK LIBRARY WEB APPLICATION

## About

This project involves the design of an API for a Book Library Web Application.
The main language used for it is Go.

## Usage

After cloning the repository, run:

```
docker-compose build
docker-compose up
```

The functionalities implemented include public and private endpoints.
The private endpoints require authentication by the users.
In this implementation, there are two users: admin (librarian) and customer.

## Endpoints for admin(librarian)

### Add a book (POST /books)

```
curl -u admin:adminpw --header "Content-Type: application/json" --request POST --data '{"Title": "my_title", "Author": "Authors Name", "Isbn": "4578", "IssueYear": 2023}' localhost:4444/books
```

### Remove a book (DELETE /books/{id})

```
curl -u admin:adminpw -X DELETE localhost:4444/books/{id}
```

### Add a customer (POST /customers)

```
curl -u admin:adminpw --header "Content-Type: application/json" --request POST --data '{"Name": "Costumers Name"}' localhost:4444/customers
```

### List all customers (GET /customers)

```
curl -u admin:adminpw localhost:4444/customers
```


## Endpoints for customer

### Borrow a book (PATCH /books/borrow/{id})

```
curl -u customer:customerpw --request PATCH localhost:4444/books/borrow/{id}
```

### Return a book (PATCH /books/return/{id})

```
curl -u customer:customerpw --request PATCH localhost:4444/books/return/{id}
```

## Public Endpoints

### List all books (GET /books)

```
curl localhost:4444/books
```

### Get book by ISBN (GET /books/{isbn})

```
curl localhost:4444/books/{isbn}
```

### Get book by ID (GET /books/{id})

```
curl localhost:4444/books/{isbn}
```

## Possible improvements

### Immediate next steps

* I tested everything manually, but with more time I would have prefered writing
automated tests for the different functionalities.

### Database migrations

In production code, the database and tables wouldn't be created as they are
right now and there would be a system in place for having migrations

### Functionality
Considering the logic of the usage for the application, these functionalities
would be relevant in future iterations of the project:

* Give admins(librarians) the possibility to keep track of which user has borrowed each book
* Increase the copies of a book instead of creating 'new books' with the same
	information. This implies a differentiation between adding a new book and
	adding a new copy of an existing book
* Have the possibility to delete customers (only after they have returned their books)
* Could add expiry times for renting books

### Validation

Adding validation for the different fields (for example to ensure that all isbns
are unique) would be helpful

### Infrastructure

In production code, an integration with a monitoring solution and CI/CD would
become relevant

### Code structure

With more time, I would have liked refactoring some of the code to have a more
clean and clear structure. For example, I would have prefered isolating the SQL queries in a
single file and making usage of it in the go files.
There's also scafolding code around validation etc, ready to be used for future
iterations

### Authentication

In future iterations and in production code there would be several admin accounts and customer
accounts, and a more robust authentication system
