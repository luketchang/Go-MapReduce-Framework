package server

import (
	"net"
	"testing"
	"time"
)

// func TestServerConnection(t *testing.T) {
// 	_, err := net.Dial("tcp", ServerAddress)
// 	if err != nil {
// 		t.Error("could not connect to server: ", err)
// 	}
// 	// defer conn.Close()
// }

func TestReceiveWorkerMessage(t *testing.T) {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()

	conn.Write([]byte(WorkerReady))
	time.Sleep(2 * time.Second)
}
