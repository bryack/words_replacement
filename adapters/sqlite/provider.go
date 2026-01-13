package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteFormsProvider struct {
	db *sql.DB
}

func NewSQLiteFormsProvider() (*SQLiteFormsProvider, error) {
	provider := &SQLiteFormsProvider{}
	err := provider.initDataBase()
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	err = provider.createTable()
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	err = provider.insertTestData()
	if err != nil {
		return nil, fmt.Errorf("failed to insert test data to db: %w", err)
	}
	return provider, nil
}

func (sfp *SQLiteFormsProvider) GetForms(word string) (singular, plural []string, err error) {
	var sing, plur string
	err = sfp.db.QueryRow("SELECT singular, plural FROM word_forms WHERE word = ?", word).Scan(&sing, &plur)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, fmt.Errorf("not found")
		} else {
			return nil, nil, fmt.Errorf("failed to scan: %w", err)
		}
	}

	singular = append(singular, sing)
	plural = append(plural, plur)

	return singular, plural, nil
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
