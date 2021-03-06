package src

import (
	"github.com/gorilla/mux"
	"maykonn/restful-storage/src/http_handlers/authorization_token"
	"maykonn/restful-storage/src/http_handlers/default_verb"
	"maykonn/restful-storage/src/http_handlers/middleware"
)

const defaultVerbsRoutes = "/{route:[\\w\\d-]{1,}}"

func Router() *mux.Router {
	MainRouter := mux.NewRouter()
	AuthorizationRouter := AuthorizationRouter()
	DefaultVerbsRouter := DefaultVerbsRouter()

	MainRouter.Handle("/authorization", AuthorizationRouter)
	MainRouter.Handle(defaultVerbsRoutes, DefaultVerbsRouter)

	return MainRouter
}

func AuthorizationRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/authorization", authorization_token.New).Methods("GET")
	return router
}

func DefaultVerbsRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(defaultVerbsRoutes, default_verb.Get).Methods("GET")
	router.HandleFunc(defaultVerbsRoutes, default_verb.Post).Methods("POST")
	router.HandleFunc(defaultVerbsRoutes, default_verb.Put).Methods("PUT")
	router.HandleFunc(defaultVerbsRoutes, default_verb.Patch).Methods("PATCH")
	router.HandleFunc(defaultVerbsRoutes, default_verb.Delete).Methods("DELETE")

	Authorization := middleware.Authorization{}
	router.Use(Authorization.Middleware)

	return router
}
