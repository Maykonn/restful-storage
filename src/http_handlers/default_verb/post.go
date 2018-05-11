package default_verb

import (
	"net/http"
	"fmt"
	"maykonn/restful-storage/src/storage"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post called")

	storage.Create(r.Body)
}
