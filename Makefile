server:
	make compile && ./build/api

compile:
	GOOS=linux go build -o ./build/api ./main.go