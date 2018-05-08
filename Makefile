install:
	go get -u github.com/codegangsta/gin
	go get -u github.com/joho/godotenv
	go get -u github.com/google/uuid
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/Maykonn/jwt-go-validation
	go get -u github.com/gorilla/mux

clean:
	rm -rf ./build/*

dev-server: clean
	gin -b ./build/api -p 8001 -a 8000 -l localhost

compile: clean
	GOOS=linux go build -o ./build/api ./main.go

start: install dev-server
deploy: install compile