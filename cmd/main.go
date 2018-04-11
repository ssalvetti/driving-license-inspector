package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/ssalvetti/driving-license-inspector/pkg/patenti"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgres://postgres:postgres@localhost/driving_licenses?sslmode=disable"
	_, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	csvFile := flag.String("csv", "", "path to csv file downloaded from gov website")
	flag.Parse()

	recordPatenti, err := patenti.ReadFromCsv(*csvFile)
	if err != nil {
		log.Fatalf("could not read csv file: %v", err)
	}
	log.Printf("records read: %d", len(recordPatenti))
}
