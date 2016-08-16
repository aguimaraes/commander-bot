package main

import (
	"log"
	"os"
)

type rtmConfig struct {
	Token string
	URL   string
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("First argument must be a token.")
	}

	config := new(rtmConfig)

	config.URL = "https://slack.com/api/rtm.start?token=%s"

	config.Token = os.Args[1]

	slack := new(slack)
	slack.startRTMSession(config)
}
