package controllers

import (
	"fmt"
	"net/http"

	"github.com/dangdennis/go-photos/backend/models"
)

// NewUsers returns our user controller
func NewUsers(us *models.UserService) *Users {
	return &Users{
		us: us,
	}
}

// Users defines our users controller
type Users struct {
	us *models.UserService
}

// SignupForm aggregates our parsed form data
type SignupForm struct {
	Name     string `schema:"name"`
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

	fmt.Println("Name :", form.Name)
	fmt.Println("Email :", form.Email)
	fmt.Println("Password :", form.Password)

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	fmt.Println("User Name :", form.Name)
	fmt.Println("User Email :", form.Email)
	fmt.Println("User Password :", form.Password)

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User is", user)
}

// curl -d "email=value1&password=value2" -X POST http://localhost:3000/signup

// LoginForm holds the values from the req body
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login is used to process the login form when a user
// tries to log in as an existing user (via email & pw).
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid password provided.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	fmt.Println("user email", user.Email)

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)
}

// CookieTest is used to display cookies set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email is:", cookie.Value)
}
