package main

import (
	"context"
	"log"

	"github.com/nataliabu/library-service/internal/database"
)

func listCustomersDB(db *database.DB, ctx context.Context) ([]database.Customer, error) {
	data := []database.Customer{}
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int32
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, database.Customer{id, name})
	}
	return data, nil
}

func addCustomerDB(db *database.DB, ctx context.Context, customer *database.Customer) (*database.Customer, error) {
	query := `INSERT INTO customers (name)
	  VALUES ($1) RETURNING id`
	var pk int32
	err := db.QueryRow(query, customer.Name).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	newCustomer := database.Customer{
		ID:   pk,
		Name: customer.Name,
	}
	return &newCustomer, nil
}
