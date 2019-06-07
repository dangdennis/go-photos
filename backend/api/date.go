package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler is the entry point for our date route
func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)

	fmt.Fprintf(w, currentTime)
}
