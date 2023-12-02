package repository_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/im11u/gql-example/go/infrastructure/database"
)

var db *sql.DB

func TestMain(m *testing.M) {
	db, _ = database.Connect()
	defer db.Close()

	os.Exit(m.Run())
}
