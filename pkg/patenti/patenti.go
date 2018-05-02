package patenti

import (
	"database/sql"
	"encoding/csv"
	"os"
)

type RecordPatente struct {
	id                  string
	anno_nascita        string
	regione_residenza   string
	provincia_residenza string
	comune_residenza    string
	sesso               string
	categoria_patente   string
	data_rilascio       string
	abilitato_a         string
	data_abilitazione_a string
	data_scadenza       string
	punti_patente       string
}

func InsertRecordPatenteToDB(db *sql.DB, rec RecordPatente) error {
	query := "INSERT INTO patenti VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	return insertToDB(db, query, rec.id,
		rec.anno_nascita, rec.regione_residenza, rec.provincia_residenza,
		rec.comune_residenza, rec.sesso, rec.categoria_patente,
		rec.data_rilascio, rec.abilitato_a, rec.data_abilitazione_a,
		rec.data_scadenza, rec.punti_patente)

}

func insertToDB(db *sql.DB, query string, args ...interface{}) error {
	if _, err := db.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func ReadFromCsv(inputFile string) ([]RecordPatente, error) {
	csvFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	listaPatenti := make([]RecordPatente, 0)
	for c, r := range records {
		if c == 0 {
			continue
		}
		record := RecordPatente{
			id:                  r[0],
			anno_nascita:        r[1],
			regione_residenza:   r[4],
			provincia_residenza: r[3],
			comune_residenza:    r[2],
			sesso:               r[6],
			categoria_patente:   r[7],
			data_rilascio:       r[8],
			abilitato_a:         r[9],
			data_abilitazione_a: r[10],
			data_scadenza:       r[11],
			punti_patente:       r[13],
		}
		listaPatenti = append(listaPatenti, record)
	}
	return listaPatenti, nil
}

func BatchInsertRecordsToDB(db *sql.DB, records []RecordPatente) error {
	query := "INSERT INTO patenti VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	for _, rec := range records {
		if _, err := stmt.Exec(rec.id,
			rec.anno_nascita, rec.regione_residenza, rec.provincia_residenza,
			rec.comune_residenza, rec.sesso, rec.categoria_patente,
			rec.data_rilascio, rec.abilitato_a, rec.data_abilitazione_a,
			rec.data_scadenza, rec.punti_patente); err != nil {
			continue
		}

	}
	return tx.Commit()
}
