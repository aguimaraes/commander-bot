package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type responseRtmStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	URL   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	ID string `json:"id"`
}

func startRTMSession(config *rtmConfig) {
	response := callRTMStart(config)
	url := parseRTMStartResponse(response)

	fmt.Printf("URL: %s", url)
}

func callRTMStart(config *rtmConfig) (response *http.Response) {

	tokenURL := fmt.Sprintf(config.URL, config.Token)

	response, err := http.Get(tokenURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected.")

	if response.StatusCode != 200 {
		log.Fatalf("HTTP Error %d - %s", response.StatusCode, response.Status)
	}

	fmt.Println("Start succeeded.")

	return
}

func parseRTMStartResponse(response *http.Response) (url string) {
	body, err := ioutil.ReadAll(response.Body)

	response.Body.Close()

	if err != nil {
		log.Fatalf("Response Parser Error: %s", err)
	}

	fmt.Println("Data received.")

	var jsonResponse responseRtmStart

	jsonErr := json.Unmarshal(body, &jsonResponse)

	if jsonErr != nil {
		log.Fatalf("JSON Parser Error:  %s", jsonErr)
	}

	fmt.Println("JSON decode succeeded.")

	if !jsonResponse.Ok {
		log.Fatalf("Slack Error: %s", jsonResponse.Error)
	}

	url = jsonResponse.URL

	fmt.Println("Oh right!")

	return
}
