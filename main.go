package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Set the content type to text/html
	fmt.Fprint(w, "<h1>Welcome to Lenslocked!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Set the content type to text/html
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:example@example.com\">example@example.com</a>.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":3000", nil)
}
