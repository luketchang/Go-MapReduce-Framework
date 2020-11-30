package server

type Worker interface {
	requestInput(name *string) bool
	processInput(name string, output string) bool
	notifyServer(name string, success bool)
	alertServerOfProgress(info string)
}

func requestInput() string {
	return ""
}
