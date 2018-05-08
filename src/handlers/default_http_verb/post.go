package default_http_verb

import (
	"net/http"
	"fmt"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post called")
}
