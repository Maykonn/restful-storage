package default_http_verb

import (
	"net/http"
	"fmt"
)

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Put called")
}
