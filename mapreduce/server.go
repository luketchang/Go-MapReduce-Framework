package mapreduce

import (
	"fmt"
	"log"
	"net"
	"os/exec"
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

func (s *Server) startServer() {
	ln, err := net.Listen("tcp", InternalServerAddress)
	if err != nil {
		log.Fatal(MapReduceError{errStartingServer, err.Error()})
	}

	log.Println("Server listening on (outside): ", ExternalServerAddress)
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

	var wg sync.WaitGroup
	wg.Add(s.numMappers)
	for i := 0; i < s.numMappers; i++ {
		mapperNode := s.getRandomNode()
		command := s.buildMapperCommand(mapperNode)
		log.Println("Starting command [", command, "] on remote command on machine:", mapperNode)
		go s.spawnWorker(command, &wg)
	}
	log.Println("Waiting for mappers to finish...")
	wg.Wait()
	log.Println("Mappers done.")
}

func (s *Server) buildMapperCommand(remoteMachine string) *exec.Cmd {
	commandFlag := fmt.Sprintf("--command=%s %s %s", s.mapper, s.mapperExecutable, s.outputPath)
	// commandFlag := fmt.Sprintf("--command=\"ls\"")
	zoneFlag := fmt.Sprintf("--zone=%s", ServerZone)
	command := exec.Command("gcloud", "compute", "ssh", remoteMachine, zoneFlag, commandFlag)
	return command
}

func (s *Server) spawnWorker(command *exec.Cmd, wg *sync.WaitGroup) {
	err := command.Start()
	if err != nil {
		log.Fatal(MapReduceError{errExecutingCmd, err.Error()})
	}
	err = command.Wait()
	log.Println("Command finished with error:", err)
	wg.Done()
}

func (s *Server) spawnReducers() {}

func (s *Server) shutDown() {
	s.serverIsRunning = false
	s.listener.Close()
}
