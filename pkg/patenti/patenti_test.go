package patenti

import (
	"database/sql"
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
	stringedIdToInsert := "347"
	query := "INSERT INTO patenti (id) VALUES($1)"
	err = insertToDB(db, query, stringedIdToInsert)
	if err != nil {
		t.Errorf("insert failed : %v", err)
	}
	var id int
	if err = db.QueryRow("SELECT id FROM patenti WHERE id=$1", 347).Scan(&id); err != nil {
		t.Errorf("select query failed %v", err)
	}
	if strconv.Itoa(id) != stringedIdToInsert {
		t.Errorf("id mismatched %d", id)
	}
}

func Test_InsertCitySingleQuote(t *testing.T) {
	testrecord := RecordPatente{
		id:                  "10",
		anno_nascita:        "1990",
		regione_residenza:   "test",
		provincia_residenza: "co",
		comune_residenza:    "cantu'",
		sesso:               "t",
		categoria_patente:   "a",
		data_rilascio:       "1990",
		abilitato_a:         "s",
		data_abilitazione_a: "1990",
		data_scadenza:       "1990",
		punti_patente:       "30",
	}
	db, err := openConnection()
	if err != nil {
		t.Errorf("connection failed: %v", err)
	}
	if err := InsertRecordPatenteToDB(db, testrecord); err != nil {
		t.Errorf("insert to DB failed: %v", err)
		t.FailNow()
	}
	var rec RecordPatente
	err = db.QueryRow("SELECT * FROM patenti WHERE id=10").Scan(&rec.id,
		&rec.anno_nascita, &rec.regione_residenza, &rec.provincia_residenza,
		&rec.comune_residenza, &rec.sesso, &rec.categoria_patente,
		&rec.data_rilascio, &rec.abilitato_a, &rec.data_abilitazione_a,
		&rec.data_scadenza, &rec.punti_patente)
	if err != nil {
		t.Errorf("convertion from db failed: %v", err)
	}
	if rec.comune_residenza != testrecord.comune_residenza {
		t.Error("comune residenza not ok")
	}
	cleanUp(db)
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

func BenchmarkInsertMultipleRecords(b *testing.B) {

	db, err := openConnection()
	if err != nil {
		b.Errorf("connection failed: %v", err)
	}
	b.N = 6000
	records := make([]RecordPatente, 0)
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
		// assign record to records list
		records = append(records, rec)
		// each 1000 i, we run the multiple insert and reset records
		if i%1000 == 0 {
			if err := BatchInsertRecordsToDB(db, records); err != nil {
				b.Errorf("failed to insert %v", err)
			}
			records = make([]RecordPatente, 0)
		}
	}
	cleanUp(db)
}

func Test_InsertMultipleRecordsInOneQuery(t *testing.T) {
	records := make([]RecordPatente, 0)
	for i := 0; i < 10; i++ {
		id := i + 1000
		testrecord := RecordPatente{
			id:                  strconv.Itoa(id),
			anno_nascita:        "1990",
			regione_residenza:   "test",
			provincia_residenza: "co",
			comune_residenza:    "cantu'",
			sesso:               "t",
			categoria_patente:   "a",
			data_rilascio:       "1990",
			abilitato_a:         "s",
			data_abilitazione_a: "1990",
			data_scadenza:       "1990",
			punti_patente:       "30",
		}
		records = append(records, testrecord)
	}
	db, err := openConnection()
	if err != nil {
		t.Errorf("connection failed: %v", err)
	}
	if err := BatchInsertRecordsToDB(db, records); err != nil {
		t.Errorf("batch failed to insert %v", err)
	}
	var count int
	if err = db.QueryRow("SELECT COUNT(*) FROM patenti").Scan(&count); err != nil {
		t.Errorf("select query failed %v", err)
	}
	if count != 10 {
		t.Errorf("expected 10 lines, having %d", count)
	}
	cleanUp(db)
}
