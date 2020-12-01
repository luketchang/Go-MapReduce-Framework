package server

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testUnprocessedList = []string{
		"00001.input",
		"00002.input",
		"00003.input",
	}
)

func TestRequestInput(t *testing.T) {
	server := NewServer([]string{})

	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	conn.Write([]byte(WorkerReady))
	serverMsg := readFromConn(conn)
	log.Println("Message received from server: ", serverMsg)
	assert.EqualValues(t, server.inflight[0], serverMsg)
}
