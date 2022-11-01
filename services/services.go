package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

var Address = []string{"https://www.saperescienza.it/news/spazio-tempo?format=feed"}

type SimpleResponse struct {
	Status string
	Value  []ItemLink
}

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

type ItemLink struct {
	Title string
	Link  string
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

func FetchTopN(address string, n int) []ItemLink {
	var TopItems []ItemLink
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return TopItems
	}
	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return TopItems
	}
	var d ItemLink
	var data Item
	for _, item := range rss.Channel.Items {
		data = item
		fmt.Printf("[%s]:[%s]\n", data.Title, data.Link)
		//Store Data Here(Establish Prior Connection With DB)
		//Store.Insert(data)
		d.Link = data.Link
		d.Title = data.Title
		TopItems = append(TopItems, d)

	}
	return TopItems[0:n]
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

func parseCommandRequest(inJson []byte) CommandRequest {

	var output CommandRequest
	json.Unmarshal(inJson, &output)
	log.Printf("\n[output:%v]\n", output)
	return output
}

func ProcessCommand(w http.ResponseWriter, req *http.Request) {
	// MLo1112o22: Receives a request in the form of a pair command,attributes
	// Processes the command and invokes the methods to handle it
	// It returns a JSON structure

	var response string

	ctx := req.Context()
	requestString, _ := ioutil.ReadAll(req.Body)
	command := parseCommandRequest(requestString)
	log.Printf("[ProcessCommand] Command:[%s] - Attribute:[%s]\n", command.Command, command.Attribute)
	switch command.Command {
	case "Fetch":
		{
			log.Println("[ProcessCommand]-Received FETCH commmand")
			retValue := FetchTopN("https://www.saperescienza.it/news/spazio-tempo?format=feed", 10)
			log.Println(retValue)
			//res, _ := fmt.Printf("{result:[%v]}", retValue)

			json.NewEncoder(w).Encode(fmt.Sprint(retValue))
		}
	default:
		{
			log.Println("[ProcessCommand]- Unrecognized command")
		}
	}
	_, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Printf("Request Received:  \n\n")
	}

	select {
	case <-time.After(1 * time.Second):
		response = "response:cacca"
		json.NewEncoder(w).Encode((response))
		//fmt.Fprintf(w, "[received command]\n")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func DefaultService(w http.ResponseWriter, req *http.Request) {
	log.Println("WebServerStarted")
	response := "Feed Reader Up and Running ..."
	json.NewEncoder(w).Encode(response)
}
