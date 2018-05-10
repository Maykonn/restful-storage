server: reflex

install: stop dependencies
	mkdir ./tmp
	make server

clear:
	-rm ./build/api

build: clear
	go build -o ./build/api ./main.go

run: build
	./build/api

reflex: stop
	reflex --start-service -r '\.go$$' make run

stop:
	echo "Stopping ./build/api if it's running"
	-kill -9 `cat ./tmp/.pid`
	-rm ./tmp/.pid

showpid:
	lsof -t -i:3000

dependencies:
	go get -u github.com/joho/godotenv
	go get -u github.com/google/uuid
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/Maykonn/jwt-go-validation
	go get -u github.com/gorilla/mux

.PHONY: build run reflex stop showpid dependencies install