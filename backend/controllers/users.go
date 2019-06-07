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

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response.")
}
