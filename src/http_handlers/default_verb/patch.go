package default_verb

import (
	"fmt"
	"net/http"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Patch called")
}
