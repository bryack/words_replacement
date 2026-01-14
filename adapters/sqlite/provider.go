package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableSQL = `CREATE TABLE IF NOT EXISTS word_forms (
        word TEXT PRIMARY KEY,
        singular TEXT,
        plural TEXT
    )`
	insertWordSQL  = `INSERT INTO word_forms (word, singular, plural) VALUES (?, ?, ?)`
	selectFormsSQL = `SELECT singular, plural FROM word_forms WHERE word = ?`
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
			s, p := entry.ExtractNominativeForms()
			if s == "" && p == "" {
				continue
			}
			_, err := stmt.Exec(entry.Word, s, p)
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
	var sing, plur string
	err = sfp.db.QueryRow(selectFormsSQL, word).Scan(&sing, &plur)
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
	query := createTableSQL
	_, err := sfp.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
