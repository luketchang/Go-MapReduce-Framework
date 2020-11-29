package server

import (
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

func receiveMessage(conn net.Conn) (Message, error) {
	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		return "", errReadingMesssage
	}

	str := string(buf)
	msg, exists := stringToMessageMap[str]

	if !exists {
		return UnknownMessage, nil
	}
	return msg, nil
}
