package main

import (
	"FeedReader/services"
	"fmt"
)

func main() {
	fmt.Println("-- .FEED READER SERVER START -- ")
	services.TestService()
	fmt.Println("-- FEED READER SERVER STOP -- ")
}
