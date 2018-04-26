package handlers

import (
	"net/http"
	"fmt"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete called")
}
