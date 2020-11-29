package server

type MapReduceError struct {
	reason string
	err    string
}

var (
	errStartingServer  = "failed to start server"
	errCouldNotConnect = "failed to establish client-server connection"
	errBadArgs         = "bad arguments"

	errEnvConfig = "error setting accessing env variables"

	errReadingMessage = "could not receive worker message"
	errWritingMessage = "could not write message to worker"
)
