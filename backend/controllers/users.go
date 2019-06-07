package controllers

import (
	"fmt"
	"net/http"
)

// Users is our controller for all user related tasks
type Users struct{}

// Create is used to process the signup form when a user tries to
// create a new user account.
func (u *Users) Create(w http.ResponseWriter, r *http.Response) {
	fmt.Fprintln(w, "This is a temporary response")
}
