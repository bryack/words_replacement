package sqlite

import (
	"testing"

	"github.com/bryack/words/specifications"
	"github.com/stretchr/testify/assert"
)

func setupSQLiteProvider(t *testing.T) *SQLiteFormsProvider {
	provider, err := NewSQLiteFormsProvider()
	assert.NoError(t, err)

	t.Cleanup(func() {
		if provider.db != nil {
			provider.db.Close()
		}
	})
	return provider
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

		err = provider.insertTestData()
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
