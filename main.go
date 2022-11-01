package main

import (
	"FeedReader/services"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-- .FEED READER SERVER START -- ")
	//define the handlers for the the services request

	http.HandleFunc("/", services.DefaultService)
	http.HandleFunc("/command", services.ProcessCommand)
	// initiate the listner
	http.ListenAndServe(":8090", nil)
	services.TestService()
	fmt.Println("-- FEED READER SERVER STOP -- ")
}
