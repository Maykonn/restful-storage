package handlers

import (
	"net/http"
	"fmt"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Patch called")
}
