package server

import (
	"log"
	"net"
)

type Worker struct{}

func (w *Worker) requestInput() (string, bool) {
	conn := w.establishConnection()
	w.sendWorkerReady(conn)

	serverResponse := readFromConn(conn)
	if isServerDoneMsg(serverResponse) {
		return "", false
	}

	return serverResponse, true
}

func (w *Worker) establishConnection() net.Conn {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		log.Fatal(errCouldNotConnect)
	}
	return conn
}

func isServerDoneMsg(serverResponse string) bool {
	return stringToMessageMap[serverResponse] == ServerDone
}
