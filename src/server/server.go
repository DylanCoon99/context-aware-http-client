package server

import (
	//"fmt"
	"log"
	"net/http"
	"github.com/DylanCoon99/context-aware-http-client/src/handlers"
)


var SERVER *http.Server


func CreateServer(port string) {


	// here we are going to instantiate a server 


	serverMux := http.NewServeMux()

	// Put handle funcs here
	serverMux.HandleFunc("GET /api/test", handlers.Get2Handler)

	SERVER = &http.Server{
		Addr: ":" + port,
		Handler: serverMux,
	}


}



func Listen() {

	log.Printf("Listening on port: %v", SERVER.Addr)
	log.Fatal(SERVER.ListenAndServe())

}