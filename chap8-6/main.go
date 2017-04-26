package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request)  {
	data := []User{
		User{
			FirstName:"San",
			LastName:"Nguyen",
			Email:"vinhsannguyen91@gmail.com",
		},
	}
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":8080", r)
}
