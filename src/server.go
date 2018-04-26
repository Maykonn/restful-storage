package src

import (
	"net/http"
	"log"
)

func StartServer() {
	err := http.ListenAndServe(":8000", Router())
	log.Fatal(err)
}
