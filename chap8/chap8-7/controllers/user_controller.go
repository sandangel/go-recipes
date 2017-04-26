package controllers

import (
	"github.com/sandangel/go-recipes/chap8/chap8-7/model"
	"net/http"
	"encoding/json"
	"github.com/sandangel/go-recipes/chap8/chap8-7/store"
	"log"
)

func GetUsers(store model.UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := store.GetUsers()
		users, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(users)
	})
}

func CreateUser(store model.UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatalf("[Controllers.CreateUser]: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = store.AddUser(user)
		if err != nil {
			if err == model.ErrorEmailExists{
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

