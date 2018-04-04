package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ssalvetti/driving-license-inspector/pkg/patenti"
)

func main() {
	connectionString := "postgres://postgres:postgres@localhost/driving_licenses?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	firstID := 6133016
	insertQuery := fmt.Sprintf("INSERT INTO patenti VALUES (%d, 1960, 'LOMBARDIA', 'LODI', 'LODI', 'F', 'B', '1979-08-22 00:00:00', 'S', '1979-08-22 00:00:00', '2019-07-21 00:00:00', 30);", firstID)
	err = patenti.InsertToDB(db, insertQuery)
	if err != nil {
		log.Fatalf("insert query failed: %v", err)
	}

}
