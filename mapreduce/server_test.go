package mapreduce

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
	server := NewServer([]string{
		"mr",
		"--mapper",
		"./mapper/mrm",
		"--reducer",
		"./reducer/mrr",
		"--config",
		"files/example.cfg",
	})

	conn, err := net.Dial("tcp", InternalServerAddress)
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	conn.Write([]byte(WorkerReady))
	serverMsg := readFromConn(conn)
	log.Println("Message received from server: ", serverMsg)
	assert.EqualValues(t, server.inflight[0], serverMsg)
}
