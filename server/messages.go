package server

import (
	"log"
	"net"
)

type Message string

const (
	WorkerReady    Message = "WORKER_READY"
	JobStarted     Message = "JOB_STARTED"
	JobInfo        Message = "JOB_INFO"
	JobSucceeded   Message = "JOB_SUCCEEDED"
	JobFailed      Message = "JOB_FAILED"
	ServerDone     Message = "SERVER_DONE"
	UnknownMessage Message = "UNKNOWN_MESSAGE"
)

var stringToMessageMap = map[string]Message{
	"WORKER_READY":  WorkerReady,
	"JOB_STARTED":   JobStarted,
	"JOB_INFO":      JobInfo,
	"JOB_SUCCEEDED": JobSucceeded,
	"JOB_FAILED":    JobFailed,
	"SERVER_DONE":   ServerDone,
}

func receiveMessage(conn net.Conn) Message {
	msgString := readMessage(conn)
	msg, exists := stringToMessageMap[msgString]

	if !exists {
		return UnknownMessage
	}
	return msg
}

func readMessage(conn net.Conn) string {
	buf := make([]byte, 1024)
	strLen, err := conn.Read(buf)
	if err != nil {
		log.Fatal(errReadingMessage)
	}
	return string(buf[:strLen])
}

func (s *Server) sendJobStart(conn net.Conn, path string) {
	_, err := conn.Write([]byte(path))
	if err != nil {
		log.Fatal(errWritingMessage)
	}
}

func (s *Server) sendServerDone(conn net.Conn) {
	_, err := conn.Write([]byte(ServerDone))
	if err != nil {
		log.Fatal(errWritingMessage)
	}
}
