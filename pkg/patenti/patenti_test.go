package patenti

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
)

func cleanUp(db *sql.DB) {
	db.Query("TRUNCATE TABLE patenti")
}

func openConnection() (*sql.DB, error) {
	connectionString := "postgres://postgres:postgres@localhost/driving_licenses?sslmode=disable"
	return sql.Open("postgres", connectionString)
}
func Test_insertWithPlaceholder(t *testing.T) {
	db, err := openConnection()
	if err != nil {
		t.Errorf("connection failed: %v", err)
	}
	cleanUp(db)
	idToInsert := 347
	query := fmt.Sprintf("INSERT INTO patenti VALUES(%d)", idToInsert)
	err = insertToDB(db, query)
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

func BenchmarkInsertRecordPatente(b *testing.B) {

	db, err := openConnection()
	if err != nil {
		b.Errorf("connection failed: %v", err)
	}
	b.N = 6000
	for i := 0; i < b.N; i++ {
		stringedID := strconv.Itoa(i + 1000)
		rec := RecordPatente{
			id:                  stringedID,
			anno_nascita:        "1998",
			regione_residenza:   "as",
			provincia_residenza: "sd",
			comune_residenza:    "df",
			sesso:               "t",
			categoria_patente:   "tes",
			data_rilascio:       "testme",
			abilitato_a:         "t",
			data_abilitazione_a: "testme",
			data_scadenza:       "testme",
			punti_patente:       "23",
		}
		if err := InsertRecordPatenteToDB(db, rec); err != nil {
			b.Errorf("failed to insert %v", err)
		}
	}
	cleanUp(db)
}
