package server

import (
	"errors"
)

var (
	errStartingServer = errors.New("failed to start server")
	errBadArgs        = errors.New("bad arguments")

	errEnvConfig = errors.New("error setting accessing env variables")

	errReadingMessage = errors.New("could not receive worker message")
	errWritingMessage = errors.New("could not write message to worker")
)
