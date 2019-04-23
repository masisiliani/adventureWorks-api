package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// inicializeRoutes
func InitRoutes() {
	router := NewRouter()
	serverErr := http.ListenAndServe(":8700", router)

	if serverErr != nil {
		log.Println("Error starting server")
		log.Println(serverErr.Error())

	} else {
		fmt.Println("Started server on - localhost:8700")
	}
}

// NewRouter
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range RoutesList {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
