package router

import (
	"net/http"
	"os"

	"github.com/cauabernardino/go-rest-api/routes"
	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GenerateRouter handles the creation of the router
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}

// GenerateLoggingRouter handles the creation of a loggin router
func GenerateLoggingRouter(r *mux.Router) http.Handler {

	loggingRouter := gorilla_handlers.LoggingHandler(os.Stdout, r)

	return loggingRouter
}
