package mapreduce

import (
	"math"
	"math/rand"
	"net"
	"strings"
	"time"
)

func (s *Server) buildIPAddrMap() {
	s.ipAddressMap = map[string]string{
		"35.236.94.23":  "machine-1",
		"34.94.186.201": "machine-2",
	}
}

func (s *Server) getNodes() {
	s.nodes = []string{
		"mapper-1",
		"mapper-2",
	}
}

func getIPAddrFromConn(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}

func (s *Server) getRandomNode() string {
	randIndex := zeroInclusiveRand(len(s.nodes))
	return s.nodes[randIndex]
}

func zeroInclusiveRand(max int) int {
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().Unix())
	return int(math.Floor(rand.Float64() * float64(max)))
}
