package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type slack struct {
	jsonResponse
}

type jsonResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	URL   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	ID string `json:"id"`
}

func (api *slack) startRTMSession(config *rtmConfig) {
	response := api.callRTMStart(config)
	url := api.parseRTMStartResponse(response)

	fmt.Printf("URL: %s", url)
}

func (api *slack) callRTMStart(config *rtmConfig) (response *http.Response) {

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

func (api *slack) parseRTMStartResponse(response *http.Response) (url string) {
	body, err := ioutil.ReadAll(response.Body)

	response.Body.Close()

	if err != nil {
		log.Fatalf("Response Parser Error: %s", err)
	}

	fmt.Println("Data received.")

	decoded := api.parseJSON(body)

	if !decoded.Ok {
		if decoded.Error == "invalid_auth" {
			log.Fatal("Authentication Failure.")
		}
		log.Fatalf("Slack Error: %s", decoded.Error)
	}

	url = decoded.URL

	fmt.Println("Oh right!")

	return
}

func (api *slack) parseJSON(body []byte) (result *jsonResponse) {

	err := json.Unmarshal(body, &result)

	if err != nil {
		log.Fatalf("JSON Parser Error: %s", err)
	}

	return result
}
