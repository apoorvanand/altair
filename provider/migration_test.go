package provider_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kodefluence/altair/entity"
	"github.com/kodefluence/altair/provider"
	"github.com/stretchr/testify/assert"
)

func TestMigrationProdiverDispatcher(t *testing.T) {
	assert.NotPanics(t, func() {
		db, _, err := sqlmock.New()
		if err != nil {
			panic(err)
		}

		dbConfig := entity.MYSQLDatabaseConfig{
			MigrationSource: "file://migration",
			Database:        "altair_development",
		}

		provider.Migration().GoMigrate(db, dbConfig)
	})
}
