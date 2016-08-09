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

func connectAndListen(token string) {

	connection := connect(token)

	if connection.StatusCode != 200 {
		log.Fatal(connection.Status)
	}

	resp := listen(connection)

	parse(resp)

}

func connect(token string) (resp *http.Response) {

	tokenURL := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)

	resp, err := http.Get(tokenURL)

	if err != nil {
		log.Fatal("Merda!")
	}

	return
}

func listen(resp *http.Response) (respObj responseRtmStart) {

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal("Merda")
	}

	err = json.Unmarshal(body, &respObj)

	if err != nil {
		log.Fatal("Merda")
	}

	return

}

func parse(resp responseRtmStart) {
	log.Fatal(resp.Error)
}
