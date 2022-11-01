package services

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

var Address = []string{"https://www.saperescienza.it/news/spazio-tempo?format=feed"}

type CommandRequest struct {
	Command   string
	Attribute string
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	City      string    `xml:"city"`
	Company   string    `xml:"company"`
	Logo      string    `xml:"logo"`
	JobType   string    `xml:"jobtype"`
	Category  string    `xml:"category"`
	PubDate   string    `xml:"date"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func TestService() {
	fmt.Println("Service Tested Successfully")
}

func Crawler(address string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}
	var data Item
	for _, item := range rss.Channel.Items {
		data = item
		fmt.Printf("[%s]:[%s]\n", data.Title, data.Link)
		//Store Data Here(Establish Prior Connection With DB)
		//Store.Insert(data)
	}
}

func WebCrawler() {
	for _, value := range Address {

		Crawler(value)
	}
}
