// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "To get in touch, please send an email "+
// 		"to <a href=\"mailto:support@lenslocked.com\">"+
// 		"support@lenslocked.com</a>.")
// }

// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>FAQ</h1>")
// }

// func error404(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Custom 404 Page</h1>")
// }

// var notFoundHandler http.Handler = http.HandlerFunc(error404)

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	r.HandleFunc("/contact", contact)
// 	r.HandleFunc("/faq", faq)
// 	r.NotFoundHandler = notFoundHandler
// 	http.ListenAndServe(":3000", r)
// }
