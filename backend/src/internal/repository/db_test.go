package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSqliteDB(t *testing.T) {
	db, err := CreateDB(&Options{
		DSN:       "sqlite://llmapi.db",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreatePostgresDB(t *testing.T) {
	db, err := CreateDB(&Options{
		DSN:       "postgres://postgres:postgres@localhost:15432/default",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")

	db, err = CreateDB(&Options{
		DSN:       "postgresql://postgres:postgres@localhost:15432/default",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreateMysqlDB(t *testing.T) {
	db, err := CreateDB(&Options{
		DSN:       "mysql://root:llmapi@tcp(localhost:13306)/llmapi?charset=utf8mb4&parseTime=True&loc=Local",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreateSqlServerDB(t *testing.T) {
	dsn := os.Getenv("SQL_SERVER_URL")
	assert.NotEmpty(t, dsn, "expected non-empty SQL_SERVER_URL environment variable")
	db, err := CreateDB(&Options{
		DSN:       dsn,
		DBLogLevel: 2,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}
