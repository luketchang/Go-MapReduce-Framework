package server

import (
	"errors"
	"log"
)

var (
	errStartingServer = errors.New("failed to start server")
	errBadArgs        = errors.New("bad arguments")

	errEnvConfig = errors.New("error setting accessing env variables")

	errReadingMesssage = errors.New("could not receive worker message")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
