.PHONY: build run reflex dependencies install stop show-pid

build:
	go build -o ./build/api ./main.go

run: build
	./build/api

reflex:
	reflex --start-service -r '\.go$$' make run

dependencies:
	go get -u github.com/joho/godotenv
	go get -u github.com/google/uuid
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/Maykonn/jwt-go-validation
	go get -u github.com/gorilla/mux

install: dependencies reflex

stop:
	echo "Stopping ./build/api if it's running"
	kill -9 `cat ./tmp/.pid`
	rm ./tmp/.pid

show-pid:
	lsof -t -i:3000
