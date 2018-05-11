server: reflex

install: stop dependencies create-tmp server

reflex: stop
	reflex --start-service -r '\.go$$' make run

run: stop build create-tmp
	./build/api

build: clear
	go build -o ./build/api ./main.go

stop: clear
	echo "Stopping ./build/api if it's running"
	-kill -9 `cat ./tmp/.pid`
	-rm ./tmp/.pid

create-tmp:
	-mkdir ./tmp

clear:
	-rm ./build/api

showpid:
	lsof -t -i:3000

dependencies:
	go get -u github.com/joho/godotenv
	go get -u github.com/go-redis/redis
	go get -u github.com/google/uuid
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/Maykonn/jwt-go-validation
	go get -u github.com/gorilla/mux

.PHONY: server install reflex run build stop create-tmp clear showpid dependencies