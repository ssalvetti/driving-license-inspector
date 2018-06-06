package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/ssalvetti/driving-license-inspector/pkg/html"
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

	//routine
	go PopulateDB(recordPatenti, db)

	http.HandleFunc("/patenti", func(w http.ResponseWriter, r *http.Request) {
		//TODO db fetches records
		// records are passed to render
		webpage, err := html.RenderRecords(recordPatenti)
		if err != nil {
			http.Error(w, "Error generating webpage", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(webpage)
	})
	log.Fatalf("could not start http server %v", http.ListenAndServe(":4444", nil))
}

func PopulateDB(recordPatenti []patenti.RecordPatente, conn *sql.DB) {
	var batchLength = 1000
	var inserted int
	batch := make([]patenti.RecordPatente, 0)
	var nextToInsert int
	for i, record := range recordPatenti {
		batch = append(batch, record)
		if i%batchLength == 0 || i == len(recordPatenti)-1 {
			err := patenti.BatchInsertRecordsToDB(conn, batch[nextToInsert:i])
			if err != nil {
				log.Printf("insert failed for batch: %v", err)
				return
			}
			nextToInsert = i + 1
			inserted += batchLength
		}
	}
	log.Printf("records read: %d", len(recordPatenti))
	log.Printf("records inserted: %d", inserted)
}
