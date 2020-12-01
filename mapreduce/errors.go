package mapreduce

type MapReduceError struct {
	reason string
	err    string
}

var (
	errStartingServer  = "failed to start server,"
	errCouldNotConnect = "failed to establish client-server connection,"
	errBadArgs         = "bad arguments,"
	errReadingInputDir = "error reading input directory files,"
	errOpeningFile     = "error opening file,"
	errScanningFile    = "error scanning file,"

	errEnvConfig = "error setting accessing env variables,"

	errReadingMessage = "could not receive worker message,"
	errWritingMessage = "could not write message to worker,"

	errBadConfigFile = "incorrect configuration file formatting,"
)
