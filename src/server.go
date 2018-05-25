package src

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	if address := os.Getenv("SERVER_ADDRESS"); address != "" {
		log.Println("Server listening at ->", address)
		panic(http.ListenAndServe(address, Router()))
	}
	panic("Invalid Http Server configuration")
}
