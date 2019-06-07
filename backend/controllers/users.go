package controllers

import (
	"fmt"
	"net/http"
)

// NewUsers returns our user controller
func NewUsers() *Users {
	return &Users{}
}

// Users defines our users controller
type Users struct {
}

// SignupForm aggregates our parsed form data
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process the signup form when a user
// tries to create a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

// curl -d "email=value1&password=value2" -X POST http://localhost:3000/signup
