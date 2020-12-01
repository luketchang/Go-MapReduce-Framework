package mapreduce

import (
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	sshKeyPath  string = "~/.ssh/mr-key"
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
	mapper            string //mr executable
	reducer           string //mrm executable
	inputPath         string
	intermediatePath  string
	outputPath        string
	mapperExecutable  string //wordcount-mapper
	reducerExecutable string //wordcount-reducer

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
		"00001.input",
		"00002.input",
		"00003.input",
	}

	s.parseArgumentList(args)
	s.nodes = s.getNodes()
	s.startServer()

	return &s
}

func (s *Server) Run() {
	s.spawnMappers()
	if !s.mapOnly {
		s.spawnReducers()
	}
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

//SCP: scp -r -i ~/.ssh/mr-key ./input/ Lukes-MacBook-Pro.local@35.236.94.23:~/input
func (s *Server) sendInputFilesToVMs() {}

func (s *Server) spawnMappers() {
	s.stageInputFiles()
	for i := 0; i < s.numMappers; i++ {
		// mapperNode := s.getRandomNode()
	}
}

//SSH: ssh -i ~/.ssh/mr-key Lukes-MacBook-Pro.local@35.236.94.23
func (s *Server) buildMapperCommand(remoteIPAddr string) string {
	// sshCommand := fmt.Sprintf("ssh -i %s %s@%s", sshKeyPath, s.host, remoteIPAddr)
	// executableCommand := strconv.Quote("mrm")
	return ""
}

func (s *Server) spawnReducers() {}

func (s *Server) shutDown() {
	s.serverIsRunning = false
	s.listener.Close()
}
