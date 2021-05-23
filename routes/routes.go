package routes

import (
	"net/http"

	"github.com/cauabernardino/go-rest-api/handlers"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var productRoutes = []Route{
	{
		URI:     "/products",
		Method:  http.MethodGet,
		Handler: handlers.GetProducts,
	},
}
