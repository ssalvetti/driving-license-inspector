package patenti

import (
	"database/sql"
	"encoding/csv"
	"os"
)

func InsertToDB(db *sql.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}

func readFromCsv(inputFile string) (int, error) {
	csvFile, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		return -1, err
	}
	return len(records), nil
}
