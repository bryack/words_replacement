package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteFormsProvider struct {
	db *sql.DB
}

func (sfp *SQLiteFormsProvider) GetForms(word string) (singular, plural []string, err error) {
	return nil, nil, fmt.Errorf("not implemented")
}

func (sfp *SQLiteFormsProvider) initDataBase() error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}

	sfp.db = db
	return nil
}

func (sfp *SQLiteFormsProvider) createTable() error {
	query := `  
   CREATE TABLE IF NOT EXISTS word_forms (  
       word TEXT,  
       singular TEXT,  
       plural TEXT  
   )`
	_, err := sfp.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func (sfp *SQLiteFormsProvider) insertTestData() error {
	query := `INSERT INTO word_forms (word, singular, plural) VALUES ('подделка', 'подделка', 'подделки')`
	_, err := sfp.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to insert into table: %w", err)
	}
	return nil
}
