package main

import (
	"FeedReader/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//v2.4.2
	logWriter, err := os.Create("feedreader.log")
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	// set the log output
	log.SetOutput(logWriter)

	fmt.Println("-- FEED READER SERVER START -- ")
	//define the handlers for the the services request

	http.HandleFunc("/", services.DefaultService)
	http.HandleFunc("/command", services.ProcessCommand)
	// initiate the listner
	http.ListenAndServe(":8090", nil)
	services.TestService()
	fmt.Println("-- FEED READER SERVER STOP -- ")
}
