package server

import (
	"fmt"
	"log"
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

	//TODO: confirmations and checks for executables
	fmt.Println("Mapper: ", s.mapper)
	fmt.Println("Reducer: ", s.reducer)
	fmt.Println("Config: ", configFileName)
}

func (s *MapReduceServer) startServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(errStartingServer)
	}

	s.listener = ln
	s.orchestrateWorkers()
}

func (s *MapReduceServer) orchestrateWorkers() {
	s.serverIsRunning = true
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(errBadArgs)
		}

		if !s.serverIsRunning {
			s.listener.Close()
		}

		log.Println("Received connection request from: ", conn.RemoteAddr())
		s.handleRequest(conn)
	}
}

func (s *MapReduceServer) handleRequest(conn net.Conn) {
	fmt.Println("Handling request...")
	msg, err := receiveMessage(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Received worker message: ", msg)
	if msg == WorkerReady {
		var path string
		s.getNextFile(&path)
	}
}

func (s *MapReduceServer) getNextFile(path *string) bool {
	if len(s.unprocessed) == 0 {
		return false
	}

	s.getAndRemoveFirstFile(path)
	s.inflight = append(s.inflight, *path)
	return true
}

func (s *MapReduceServer) getAndRemoveFirstFile(path *string) {
	*path = s.unprocessed[0]
	s.unprocessed = s.unprocessed[1:]
}
