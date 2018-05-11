package src

import (
	"github.com/gorilla/mux"
	"maykonn/restful-storage/src/handlers/authorization_token"
	"maykonn/restful-storage/src/handlers/default_http_verb"
	"maykonn/restful-storage/src/handlers/middleware"
)

const defaultVerbsRoutes = "/{route:[\\w\\d-]{1,}}"

func Router() *mux.Router {
	MainRouter := mux.NewRouter()
	AuthorizationRouter := AuthorizationRouter()
	DefaultVerbsRouter := DefaultVerbsRouter()

	MainRouter.Handle("/authorization", AuthorizationRouter)
	MainRouter.Handle(defaultVerbsRoutes, DefaultVerbsRouter)

	return MainRouter
}

func AuthorizationRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/authorization", authorization_token.New).Methods("GET")
	return router
}

func DefaultVerbsRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(defaultVerbsRoutes, default_http_verb.Get).Methods("GET")
	router.HandleFunc(defaultVerbsRoutes, default_http_verb.Post).Methods("POST")
	router.HandleFunc(defaultVerbsRoutes, default_http_verb.Put).Methods("PUT")
	router.HandleFunc(defaultVerbsRoutes, default_http_verb.Patch).Methods("PATCH")
	router.HandleFunc(defaultVerbsRoutes, default_http_verb.Delete).Methods("DELETE")

	Authorization := middleware.Authorization{}
	router.Use(Authorization.Middleware)

	return router
}



/*
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &person)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

 */