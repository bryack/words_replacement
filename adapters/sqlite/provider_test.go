package sqlite

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/bryack/words/specifications"
	"github.com/stretchr/testify/assert"
)

func setupSQLiteProvider(t *testing.T) *SQLiteFormsProvider {
	t.Helper()
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "testDB.db")
	provider, err := NewSQLiteFormsProvider(dbPath, insertTestData)
	if err != nil {
		t.Fatalf("Failed to create provider: %v", err)
	}

	t.Cleanup(func() {
		if provider != nil && provider.db != nil {
			provider.db.Close()
		}
	})
	return provider
}

func insertTestData(sfp *SQLiteFormsProvider) error {
	query := `INSERT INTO word_forms (word, singular_forms, plural_forms) VALUES ('подделка', '["подделка","подделку","подделки","подделке","подделкой"]', '["подделки","подделок","подделкам","подделками","подделках"]')`
	_, err := sfp.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to insert into table: %w", err)
	}
	return nil
}

func TestSQLiteProvider(t *testing.T) {

	t.Run("specification test", func(t *testing.T) {
		provider := setupSQLiteProvider(t)
		specifications.WiktionaryFormsSpecification(t, provider)
	})
}

func TestSQLiteFormsProvider(t *testing.T) {

	t.Run("unit test for sqlite provider creation", func(t *testing.T) {
		provider := &SQLiteFormsProvider{}
		assert.NotNil(t, provider)
	})

	t.Run("unit test for sqlite provider setup", func(t *testing.T) {
		provider := setupSQLiteProvider(t)
		assert.NotNil(t, provider)
		assert.NotNil(t, provider.db)
	})
}

func TestGetForms(t *testing.T) {

	t.Run("returns forms for existing word", func(t *testing.T) {
		provider := setupSQLiteProvider(t)
		s, p, err := provider.GetForms("подделка")
		assert.NoError(t, err)
		assert.Contains(t, s, "подделка")
		assert.Contains(t, p, "подделки")
	})
	t.Run("returns error for nonexisting word", func(t *testing.T) {
		provider := setupSQLiteProvider(t)
		_, _, err := provider.GetForms("несуществующееслово")
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}

func TestLoadFromJSONLFile(t *testing.T) {
	t.Run("loads data from JSONL file into database", func(t *testing.T) {
		tempDir := t.TempDir()
		dbPath := filepath.Join(tempDir, "testDB.db")
		filepath := "fake.jsonl"
		loader := LoadFromJSONLFile(filepath)
		provider, err := NewSQLiteFormsProvider(dbPath, loader)
		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}

		var count int
		var singular, plural string
		err = provider.db.QueryRow("SELECT COUNT(*), singular_forms, plural_forms FROM word_forms WHERE word = 'подделка'").Scan(&count, &singular, &plural)
		assert.NoError(t, err)
		assert.True(t, count > 0)
		assert.Equal(t, `["подделка","подделку","подделки","подделке","подделкой","подделкою"]`, singular)
		assert.Equal(t, `["подделки","подделок","подделкам","подделками","подделках"]`, plural)
	})
}

func TestDatabasePersistence(t *testing.T) {
	t.Run("acceptance test that verifies database persistence across provider instances", func(t *testing.T) {
		tempDir := t.TempDir()
		dbPath := filepath.Join(tempDir, "testDB.db")
		provider, err := NewSQLiteFormsProvider(dbPath, insertTestData)
		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}

		s, _, err := provider.GetForms("подделка")
		assert.NoError(t, err)
		assert.Contains(t, s, "подделка")

		t.Cleanup(func() {
			if provider != nil && provider.db != nil {
				provider.db.Close()
			}
		})

		provider2, err := NewSQLiteFormsProvider(dbPath, insertTestData)
		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}

		s2, _, err := provider2.GetForms("подделка")
		assert.NoError(t, err)
		assert.Contains(t, s2, "подделка")
		assert.Equal(t, s, s2)

		t.Cleanup(func() {
			if provider != nil && provider2.db != nil {
				provider2.db.Close()
			}
		})

	})
}
