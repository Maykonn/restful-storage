package default_http_verb

import (
	"net/http"
	"fmt"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get called")
}