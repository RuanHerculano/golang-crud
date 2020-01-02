package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type User struct {
    Name   string `json:"name"`
    Active bool   `json:"active"`
}
 
type Users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleFunc).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}

func HandleFunc(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	
    users := Users{
        User{Name: "Write presentation", Active: false},
        User{Name: "Host meetup", Active: false},
    }

	json.NewEncoder(responseWriter).Encode(users);
}
