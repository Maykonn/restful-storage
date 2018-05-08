package src

import (
	"github.com/gorilla/mux"
	"maykonn/api/src/handlers/default_http_verb"
	"maykonn/api/src/handlers/authorization_token"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/authorization", authorization_token.New).Methods("GET")

	const defaultRoute = "/{route:[\\w\\d-_]{1,}}"
	router.HandleFunc(defaultRoute, default_http_verb.Get).Methods("GET")
	router.HandleFunc(defaultRoute, default_http_verb.Post).Methods("POST")
	router.HandleFunc(defaultRoute, default_http_verb.Put).Methods("PUT")
	router.HandleFunc(defaultRoute, default_http_verb.Patch).Methods("PATCH")
	router.HandleFunc(defaultRoute, default_http_verb.Delete).Methods("DELETE")

	http.Handle("/", router)

	return router
}
