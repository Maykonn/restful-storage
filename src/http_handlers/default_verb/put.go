package default_verb

import (
	"fmt"
	"net/http"
)

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Put called")
}
