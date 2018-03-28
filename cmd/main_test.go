package main

import (
	"database/sql"
	"fmt"
	"testing"
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
