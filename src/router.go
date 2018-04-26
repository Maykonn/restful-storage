package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler).Methods("GET")
	return router
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}