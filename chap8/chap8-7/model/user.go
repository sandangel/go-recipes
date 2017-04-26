package model

import "errors"

var ErrorEmailExists = errors.New("Email ID is exists")

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserStore interface {
	GetUsers() []User
	AddUser(user User) error
}
