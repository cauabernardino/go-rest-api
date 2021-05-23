package router

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/cauabernardino/go-rest-api/routes"
	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GenerateRouter handles the creation of the router
func GenerateRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r, db)
}

// GenerateLoggingRouter handles the creation of a logging router
func GenerateLoggingRouter(r *mux.Router) http.Handler {

	loggingRouter := gorilla_handlers.LoggingHandler(os.Stdout, r)

	return loggingRouter
}
