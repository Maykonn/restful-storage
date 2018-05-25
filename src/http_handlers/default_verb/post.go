package default_verb

import (
	"fmt"
	"maykonn/restful-storage/src/storage"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post called")

	storage.Create(r.Body)
}
