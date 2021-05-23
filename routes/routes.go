package routes

import (
	"net/http"

	"github.com/cauabernardino/go-rest-api/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Configure handles the configuration of the endpoints of the API
func Configure(r *mux.Router) *mux.Router {
	routes := productRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}

var productRoutes = []Route{
	{
		URI:     "/products",
		Method:  http.MethodGet,
		Handler: handlers.GetProducts,
	},
}
