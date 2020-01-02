package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type User struct {
    Name   string `json:"name"`
    Active bool   `json:"active"`
}
 
type Users []User

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleFunc).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
    users := Users{
        User{Name: "Write presentation", Active: false},
        User{Name: "Host meetup", Active: false},
    }

	json.NewEncoder(w).Encode(users);
}
