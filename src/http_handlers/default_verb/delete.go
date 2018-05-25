package default_verb

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete called")
}
