package handlers


import (
	"log"
	"fmt"
	"html"
	"net/http"
)


func GetHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Hey I got the request! Here is the method: %v", r.Method)
}


func Get2Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
