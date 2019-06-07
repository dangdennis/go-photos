package main

import (
	"net/http"

	"github.com/dangdennis/google-photos/backend/controllers"

	"github.com/gorilla/mux"
)

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
