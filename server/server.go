package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
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

func init() {
	s := Server{}
	go s.startServer()
}

func (s *Server) Run(args []string) {
	s.host = GetHost()
	s.cwd = GetCwd()
	s.verbose = true
	s.mapOnly = false
	s.serverIsRunning = false

	s.parseArgumentList(args)
	s.nodes = getNodes()
	s.startServer()
}

func (s *Server) parseArgumentList(args []string) {
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

func (s *Server) startServer() {
	ln, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatal(errStartingServer)
	}

	s.listener = ln
	s.orchestrateWorkers()
}

func (s *Server) orchestrateWorkers() {
	s.serverIsRunning = true
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(errBadArgs)
		}
		defer conn.Close()

		if !s.serverIsRunning {
			s.listener.Close()
		}

		log.Println("Received connection request from: ", conn.RemoteAddr())
		s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	fmt.Println("Handling request...")
	msg := receiveMessage(conn)

	log.Println("Received worker message: ", msg)
	if msg == WorkerReady {
		var path string
		s.fileLock.Lock()
		success := s.getNextFile(&path)
		s.fileLock.Unlock()
		if success {
			s.sendJobStart(conn, path)
		}
	}
}
