package server

import (
	"fmt"
	"net"
	"sync"
)

type MapReduceServer struct {
	host string
	cwd  string

	listener          net.Listener //equivalent of serverSocket
	verbose           bool
	mapOnly           bool
	numMappers        int
	numReducers       int
	mapper            string //from parsed cli argument
	reducer           string //from parsed cli argument
	inputPath         string
	intermediatePath  string
	outputPath        string
	mapperExecutable  string //from config file
	reducerExecutable string //from config file

	nodes           []string
	ipAddressMap    map[string]string
	serverIsRunning bool
	//no way to store thread?

	unprocessed []string
	inflight    []string
	fileLock    *sync.Mutex
}

const (
	mapperFlag  string = "--mapper"
	reducerFlag string = "--reducer"
	configFlag  string = "--config"
)

func (s *MapReduceServer) NewServer(args []string) {
	s.host = GetHost()
	s.cwd = GetCwd()
	s.verbose = true
	s.mapOnly = false
	s.serverIsRunning = false

	s.parseArgumentList(args)
	s.nodes = getNodes()
	s.startServer()
}

func (s *MapReduceServer) parseArgumentList(args []string) {
	var configFileName string
	for i := 0; i < len(args); i++ {
		ch := args[i]

		//TODO: convert to using getopt package + error checking
		if ch == mapperFlag {
			s.mapper = args[i+1]
		}
		if ch == reducerFlag {
			s.reducer = args[i+1]
		}
		if ch == configFlag {
			configFileName = args[i+1]
		}
	}

	fmt.Println("Mapper: ", s.mapper)
	fmt.Println("Reducer: ", s.reducer)
	fmt.Println("Config: ", configFileName)

	//TODO: confirmations and checks for executables
}

func (s *MapReduceServer) startServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	check(err)

	s.listener = ln
	go s.orchestrateWorkers()
}

func (s *MapReduceServer) orchestrateWorkers() {
	s.serverIsRunning = true
	for {
		conn, err := s.listener.Accept()
		check(err)

		if !s.serverIsRunning {
			s.listener.Close()
		}

		fmt.Println("Received connection request from: ", conn.RemoteAddr())
		s.handleRequest(conn)
	}
}

func (s *MapReduceServer) handleRequest(conn net.Conn) {
	fmt.Println("Handling request...")
}
