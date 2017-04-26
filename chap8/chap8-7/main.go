package main

import (
	"github.com/gorilla/mux"
	"github.com/sandangel/go-recipes/chap8/chap8-7/store"
	"github.com/sandangel/go-recipes/chap8/chap8-7/controllers"
	"net/http"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	userStore := &store.MongoUserStore{}
	r.Handle("/users", controllers.CreateUser(userStore)).Methods("POST")
	r.Handle("/users", controllers.GetUsers(userStore)).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", setUserRoutes())
}