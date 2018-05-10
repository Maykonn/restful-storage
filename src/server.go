package src

import (
	"os"
	"log"
	"net/http"
)

func StartServer() {
	if address := os.Getenv("SERVER_ADDRESS"); address != "" {
		if port := os.Getenv("SERVER_PORT"); port != "" {
			address += ":" + port
			log.Println("Server listening at ->", address)
			panic(http.ListenAndServe(address, Router()))
		}
	}
	panic("Server configuration is invalid")
}
