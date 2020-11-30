package server

import (
	"log"
	"net"
)

// type Worker interface {
// 	requestInput(name *string) bool
// 	processInput(name string, output string) bool
// 	notifyServer(name string, success bool)
// 	alertServerOfProgress(info string)
// }

type Worker struct {
}

func (w *Worker) requestInput() (string, bool) {
	conn := establishConnection()
	w.sendWorkerReady(conn)

	serverResponse := readFromConn(conn)
	if isServerDoneMsg(serverResponse) {
		return "", false
	}

	return serverResponse, true
}

func establishConnection() net.Conn {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		log.Fatal(errCouldNotConnect)
	}
	return conn
}

func isServerDoneMsg(serverResponse string) bool {
	return stringToMessageMap[serverResponse] == ServerDone
}
