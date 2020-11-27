package server

import (
	"net"
	"sync"

	mrenv "github.com/mapreduce/env"
)

type MapReduceServer struct {
	host string
	cwd  string

	listener          *net.Listener //equivalent of serverSocket
	verbose           bool
	mapOnly           bool
	numMappers        int
	numReducers       int
	mapper            string //from parsed cli argument
	reducer           string //from parsed cli argument
	inputPath         string
	intermediatePath  string
	outputPath        string
	mapperExecutable  string
	reducerExecutable string

	nodes           []string
	ipAddressMap    map[string]string
	serverIsRunning bool
	//no way to store thread?

	unprocessed []string
	inflight    []string
	fileLock    *sync.Mutex
}

func (s *MapReduceServer) newServer() {
	s.host = mrenv.GetHost()
	s.cwd = mrenv.GetCwd()
	s.verbose = true
	s.mapOnly = false
	s.serverIsRunning = false

	//s.parseArgumentList(...)
}
