package client

import (
	"fmt"
	"net/http"
)




func CreateClient() {
	fmt.Println("Hello, I am a client!")
}


func SendRequest(url string, method string, payload string) {


	switch method {
	case "GET":
		// handle GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println(resp.StatusCode)
	case "POST":
		// handle POST request
		// To be implemented
	default:
		fmt.Println("Sorry, you can only send GET and POST requests.")
	}



}