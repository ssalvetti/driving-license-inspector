package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgres://postgres:postgres@localhost/driving_licenses?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	_, err = db.Query("INSERT INTO patenti VALUES (6133016, 1960, 'LOMBARDIA', 'LODI', 'LODI', 'F', 'B', '1979-08-22 00:00:00', 'S', '1979-08-22 00:00:00', '2019-07-21 00:00:00', 30);")
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
}
