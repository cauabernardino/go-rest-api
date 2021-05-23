package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/cauabernardino/go-rest-api/config"
)

// testDB is the connection to the testa database
var testDB *sql.DB

func TestMain(m *testing.M) {

	config.LoadEnvs("testPackage")

	var err error
	testDB, err = Connect()
	if err != nil {
		log.Fatal("could not connect to test database")
	}
	defer testDB.Close()

	os.Exit(m.Run())
}
