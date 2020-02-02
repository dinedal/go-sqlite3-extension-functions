package gosqlite3_extension_functions

import (
	"database/sql"
	"testing"
)

func TestOpenReturnsWithoutError(t *testing.T) {
	db, err := sql.Open("sqlite3-extension-functions", ":memory:")

	if err != nil {
		t.Fatalf(err.Error())
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf(err.Error())
	}
}
