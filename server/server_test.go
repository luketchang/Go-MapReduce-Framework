package server

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testUnprocessedList = []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}
)

func configServer() *Server {
	s := Server{}
	s.host = GetHost()
	s.cwd = GetCwd()
	s.verbose = true
	s.mapOnly = false
	s.serverIsRunning = false
	s.unprocessed = testUnprocessedList
	return &s
}

func TestRequestInput(t *testing.T) {
	server := configServer()
	server.startServer()

	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	conn.Write([]byte(WorkerReady))
	serverMsg := getMessageString(conn)
	log.Println("Message received from server: ", serverMsg)
	assert.EqualValues(t, testUnprocessedList[0], serverMsg)
}
