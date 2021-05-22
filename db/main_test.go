package db

import (
	"os"
	"testing"

	"github.com/cauabernardino/go-rest-api/config"
)

func TestMain(m *testing.M) {

	config.LoadEnvs("test")

	os.Exit(m.Run())
}
