package handlers


import (
	"log"
	"fmt"
	"html"
	"net/http"
)


func GetHandler(w http.ResponseWriter, r *http.Request) {

	log.Print("Hey I got the request!")
}


func Get2Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
