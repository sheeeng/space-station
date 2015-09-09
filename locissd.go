package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

/*
Example of JSON Response:
{
  "iss_position": {
    "latitude": -37.574653303613516,
    "longitude": 24.333194371046993
  },
  "message": "success",
  "timestamp": 1441726350
}
*/

// Use nested structs in Go to match the nested structure in JSON
// http://stackoverflow.com/a/25966781/4763512

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type IssNow struct {
	Message   string  `json:"message"`
	Timestamp uint32  `json:"timestamp"`
	Position  LatLong `json:"iss_position"`
}

func main() {
	var body []byte
	var response *http.Response
	var request *http.Request

	url := "http://api.open-notify.org/iss-now.json"
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

	var data IssNow
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", data)
	fmt.Printf("Time: %v | Lat./Long. : %v, %v\n", data.Timestamp, data.Position.Latitude, data.Position.Longitude)
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("Error: %s\n\n", err)
	}
}
