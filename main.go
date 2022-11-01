package main

import (
	"FeedReader/services"
	"fmt"
)

func main() {
	fmt.Println("-- .FEED READER SERVER START -- ")
	services.WebCrawler()
	services.TestService()
	fmt.Println("-- FEED READER SERVER STOP -- ")
}
