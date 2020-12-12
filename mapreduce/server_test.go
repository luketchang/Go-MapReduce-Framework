package mapreduce

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestInputWorkerReady(t *testing.T) {
	s := Server{}
	s.address = ":8001"
	s.unprocessed = []string{
		"test_files/input/00001.input",
		"test_files/input/00002.input",
		"test_files/input/00003.input"}
	s.inflight = make(map[string]bool)
	s.startServer()

	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	sendMessage(conn, WorkerReady.toString(), "")
	serverMsg := readFromConn(conn)
	assert.EqualValues(t, s.inflight[serverMsg], true)
}

func TestRequestInputNoMoreFiles(t *testing.T) {
	s := Server{}
	s.address = ":8002"
	s.unprocessed = []string{}
	s.startServer()

	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	sendMessage(conn, WorkerReady.toString(), "")
	serverMsg := readFromConn(conn)
	assert.EqualValues(t, serverMsg, ServerDone)
}
