package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableSQL = `CREATE TABLE IF NOT EXISTS word_forms (
        word TEXT PRIMARY KEY,
        singular_forms TEXT, -- JSON array: ["подделка", "подделку"]
        plural_forms TEXT -- JSON array: ["подделки"]
    )`
	insertWordSQL  = `INSERT INTO word_forms (word, singular_forms, plural_forms) VALUES (?, ?, ?)`
	selectFormsSQL = `SELECT singular_forms, plural_forms FROM word_forms WHERE word = ?`
)

type SQLiteFormsProvider struct {
	db *sql.DB
}

// DataLoader defines a function that loads word forms into the database.
type DataLoader func(sfp *SQLiteFormsProvider) error

func NewSQLiteFormsProvider(loader DataLoader) (*SQLiteFormsProvider, error) {
	provider := &SQLiteFormsProvider{}
	err := provider.initDataBase()
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	err = provider.createTable()
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	err = loader(provider)
	if err != nil {
		return nil, fmt.Errorf("failed to load data to db: %w", err)
	}
	return provider, nil
}

func LoadFromJSONLFile(filepath string) DataLoader {
	return func(sfp *SQLiteFormsProvider) error {
		entries, err := LoadFromJSONL(filepath)
		if err != nil {
			return fmt.Errorf("failed to load from JSONL file %q: %w", filepath, err)
		}
		tx, err := sfp.db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare(insertWordSQL)
		if err != nil {
			return fmt.Errorf("failed to prepare statement: %w", err)
		}
		defer stmt.Close()

		for _, entry := range entries {
			s, p := entry.ExtractTwoForms()
			if len(s) == 0 && len(p) == 0 {
				continue
			}
			singularJSON, err := json.Marshal(s)
			if err != nil {
				return fmt.Errorf("failed to marshal %+v: %w", s, err)
			}
			pluralJSON, err := json.Marshal(p)
			if err != nil {
				return fmt.Errorf("failed to marshal %+v: %w", p, err)
			}
			_, err = stmt.Exec(entry.Word, string(singularJSON), string(pluralJSON))
			if err != nil {
				return fmt.Errorf("failed to insert word %q: %w", entry.Word, err)
			}
		}
		if err = tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
		return nil
	}
}

func (sfp *SQLiteFormsProvider) GetForms(word string) (singular, plural []string, err error) {
	var singularJSON, pluralJSON string
	err = sfp.db.QueryRow(selectFormsSQL, word).Scan(&singularJSON, &pluralJSON)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to scan: %w", err)
	}

	if err := json.Unmarshal([]byte(singularJSON), &singular); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal %q: %w", singularJSON, err)
	}
	if err := json.Unmarshal([]byte(pluralJSON), &plural); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal %q: %w", pluralJSON, err)
	}

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
	query := createTableSQL
	_, err := sfp.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
