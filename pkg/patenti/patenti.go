package patenti

import "database/sql"

func InsertToDB(db *sql.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
