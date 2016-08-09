package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("First argument must be a token.")
	}

	token := os.Args[1]

	connectAndListen(token)
}
