package storage

import "time"

func Create(data interface{}) {
	// implements cache...
	Redis().Set("document.Id", "baita", 86400*time.Second)

	// implements async queued persistence...
}
