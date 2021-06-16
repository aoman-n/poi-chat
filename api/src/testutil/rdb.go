package testutil

import (
	"testing"

	"github.com/laster18/poi/api/src/infra/db"
	"gorm.io/gorm"
)

func SetupRDB(t *testing.T) *gorm.DB {
	t.Helper()
	db := db.NewDb()
	tx := db.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})
	return tx
}
