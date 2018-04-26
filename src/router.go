package src

import (
	"github.com/gorilla/mux"
	"maykonn/api/src/handlers"
)

func Router() *mux.Router {
	const route = "/[\\w\\d-_]{64}"

	router := mux.NewRouter()
	router.HandleFunc(route, handlers.Get).Methods("GET")
	router.HandleFunc(route, handlers.Post).Methods("POST")
	router.HandleFunc(route, handlers.Put).Methods("PUT")
	router.HandleFunc(route, handlers.Delete).Methods("DELETE")

	return router
}
