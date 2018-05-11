package storage

func Create(data interface{}) {
	// implements cache...
	Redis().Set("document.Id", data, 86400)

	// implements async queued persistence...
}