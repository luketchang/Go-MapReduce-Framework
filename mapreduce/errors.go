package mapreduce

type MapReduceError struct {
	reason string
	err    string
}

const (
	WrongNumArgs    int = 2
	InvalidArgsType int = 3
)

var ErrorCodeToStringMap = map[int]string{
	WrongNumArgs:    "Invalid number of arguments.",
	InvalidArgsType: "Invalid argument(s) type.",
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

	errExecutingCmd = "failed to execute command on remote machine,"
)
