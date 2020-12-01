package server

import (
	"log"
	"net"
	"sync"
)

const (
	mapperFlag  string = "--mapper"
	reducerFlag string = "--reducer"
	configFlag  string = "--config"
)

type Server struct {
	host string
	cwd  string

	listener          net.Listener
	verbose           bool
	mapOnly           bool
	numMappers        int
	numReducers       int
	mapper            string
	reducer           string
	inputPath         string
	intermediatePath  string
	outputPath        string
	mapperExecutable  string
	reducerExecutable string

	nodes           []string
	ipAddressMap    map[string]string
	serverIsRunning bool

	unprocessed []string
	inflight    []string
	fileLock    sync.Mutex
}

func NewServer(args []string) *Server {
	s := Server{}
	s.host = GetHost()
	s.cwd = GetCwd()
	s.verbose = true
	s.mapOnly = false
	s.serverIsRunning = false
	s.unprocessed = []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	s.parseArgumentList(args)
	s.nodes = getNodes()
	s.startServer()

	return &s
}

func (s *Server) Run() {
	s.spawnMappers()
	if !s.mapOnly {
		s.spawnReducers()
	}
}

func (s *Server) startServer() {
	ln, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatal(MapReduceError{errStartingServer, err.Error()})
	}

	log.Println("Server listening on: ", ServerAddress)
	s.listener = ln
	go s.orchestrateWorkers()
}

func (s *Server) orchestrateWorkers() {
	s.serverIsRunning = true
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(MapReduceError{errCouldNotConnect, err.Error()})
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
	msg := receiveMessage(conn)

	log.Println("Received worker message: ", msg)
	if msg == WorkerReady {
		s.fileLock.Lock()
		path := s.getNextFile()
		s.fileLock.Unlock()
		if !isEmpty(path) {
			s.sendJobStart(conn, path)
		} else {
			s.sendServerDone(conn)
		}
		//TODO: add other message cases after worker code
	}
}

func (s *Server) spawnMappers() {
	s.stageInputFiles()
	for i := 0; i < s.numMappers; i++ {
		// mapperNode := s.getRandomNode()
	}
}

func (s *Server) spawnReducers() {

}

func (s *Server) shutDown() {
	s.serverIsRunning = false
	s.listener.Close()
}
