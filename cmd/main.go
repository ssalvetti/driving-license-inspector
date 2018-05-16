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
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	csvFile := flag.String("csv", "", "path to csv file downloaded from gov website")
	flag.Parse()

	recordPatenti, err := patenti.ReadFromCsv(*csvFile)
	if err != nil {
		log.Fatalf("could not read csv file: %v", err)
	}
	var batchLength = 1000
	var inserted int
	batch := make([]patenti.RecordPatente, 0)
	var nextToInsert int
	for i, record := range recordPatenti {
		batch = append(batch, record)
		if i%batchLength == 0 {
			err := patenti.BatchInsertRecordsToDB(db, batch[nextToInsert:i])
			if err != nil {
				log.Printf("insert failed for batch: %v", err)
				continue
			}
			nextToInsert = i + 1
			inserted += batchLength
		}
	}

	log.Printf("records read: %d", len(recordPatenti))
	log.Printf("records inserted: %d", inserted)
}
