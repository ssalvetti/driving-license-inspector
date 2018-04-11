package patenti

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
)

func cleanUp(db *sql.DB) {
	db.Query("TRUNCATE TABLE patenti")
}
func Test_insertWithPlaceholder(t *testing.T) {
	connectionString := "postgres://postgres:postgres@localhost/driving_licenses?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Errorf("connection failed: %v", err)
	}
	cleanUp(db)
	idToInsert := 347
	query := fmt.Sprintf("INSERT INTO patenti VALUES(%d)", idToInsert)
	err = InsertToDB(db, query)
	if err != nil {
		t.Errorf("insert failed : %v", err)
	}
	var id int
	if err = db.QueryRow(fmt.Sprintf("SELECT id FROM patenti WHERE id=%d", idToInsert)).Scan(&id); err != nil {
		t.Errorf("select query failed %v", err)
	}
	if id != idToInsert {
		t.Errorf("id mismatched %d", id)
	}
}

func Test_ReadCsvFile(t *testing.T) {
	testFilePath := filepath.Join("..", "..", "test", "fixtures", "lombardia-subset.csv")
	records, err := ReadFromCsv(testFilePath)
	if err != nil {
		t.Errorf("reading from csv failed: %v", err)
	}
	if len(records) <= 0 {
		t.Error("no lines read")
	}
}

func TestReadCsvIntoCustomStructure(t *testing.T) {
	testFilePath := filepath.Join("..", "..", "test", "fixtures", "lombardia-subset.csv")
	records, err := ReadFromCsv(testFilePath)
	if err != nil {
		t.Errorf("reading from csv failed: %v", err)
	}
	if records[0].id != "6133015" {
		t.Error("unexpected value")
	}
	if records[1].categoria_patente != "B" {
		t.Error("wrong category")
	}
}
