package sqlite

import (
	"fmt"
	"testing"

	"github.com/bryack/words/specifications"
	"github.com/stretchr/testify/assert"
)

func setupSQLiteProvider(t *testing.T) *SQLiteFormsProvider {
	t.Helper()
	provider, err := NewSQLiteFormsProvider(insertTestData)
	assert.NoError(t, err)

	t.Cleanup(func() {
		if provider.db != nil {
			provider.db.Close()
		}
	})
	return provider
}

func insertTestData(sfp *SQLiteFormsProvider) error {
	query := `INSERT INTO word_forms (word, singular, plural) VALUES ('подделка', 'подделка', 'подделки')`
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

func TestDataBaseConnection(t *testing.T) {
	t.Run("DB insert", func(t *testing.T) {
		provider := &SQLiteFormsProvider{}
		err := provider.initDataBase()
		assert.NoError(t, err)

		err = provider.createTable()
		assert.NoError(t, err)

		err = insertTestData(provider)
		assert.NoError(t, err)

		var count int
		var plural string
		err = provider.db.QueryRow("SELECT COUNT(*), plural FROM word_forms WHERE word = 'подделка'").Scan(&count, &plural)
		assert.NoError(t, err)
		assert.True(t, count > 0)
		assert.Equal(t, "подделки", plural)
	})
}

func TestGetForms(t *testing.T) {

	t.Run("TestCase", func(t *testing.T) {
		provider := setupSQLiteProvider(t)
		s, p, err := provider.GetForms("подделка")
		assert.NoError(t, err)
		assert.Contains(t, s, "подделка")
		assert.Contains(t, p, "подделки")
	})
}

func TestLoadFromJSONLFile(t *testing.T) {
	t.Run("loads data from JSONL file into database", func(t *testing.T) {
		filepath := "fake.jsonl"
		loader := LoadFromJSONLFile(filepath)
		provider, err := NewSQLiteFormsProvider(loader)
		assert.NoError(t, err)

		var count int
		var singular, plural string
		err = provider.db.QueryRow("SELECT COUNT(*), singular, plural FROM word_forms WHERE word = 'подделка'").Scan(&count, &singular, &plural)
		assert.NoError(t, err)
		assert.True(t, count > 0)
		assert.Equal(t, "подделка", singular)
		assert.Equal(t, "подделки", plural)
	})
}
