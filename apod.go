package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// NASA Open APIs
	// https://api.nasa.gov/api.html#authentication
	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&format=JSON")
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	fmt.Printf("%s", body)
}
