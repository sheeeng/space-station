package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"
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

func locateiss() {
	response, err := http.Get("http://api.open-notify.org/iss-now.json")
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	// fmt.Printf("%s\n", body)

	var data IssNow
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", data)
	fmt.Printf("Time: %v | Lat./Long. : %v, %v\n", data.Timestamp, data.Position.Latitude, data.Position.Longitude)
}

func main() {
	// Capture ctrl+c and stop CPU profiler.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Captured %v, stopping profiler & exiting...\n",
				sig)
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

	for {
		locateiss()
		time.Sleep(time.Second * 5)
	}
}
