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

func InsertToDB(db *sql.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
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
