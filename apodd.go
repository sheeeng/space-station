package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var body []byte
	var response *http.Response
	var request *http.Request

	// NASA Open APIs
	// https://api.nasa.gov/api.html#authentication
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&format=JSON"
	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		debug(httputil.DumpRequestOut(request, true))
		response, err = (&http.Client{}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		debug(httputil.DumpResponse(response, true))
		body, err = ioutil.ReadAll(response.Body)
	}

	if err == nil {
		fmt.Printf("%s\n\n", body)
	} else {
		log.Fatalf("Error: %s\n\n", err)
	}
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("Error: %s\n\n", err)
	}
}
