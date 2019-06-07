package handler

import (
	"fmt"
	"net/http"
)

// Handler is the entry point for our signup route
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
