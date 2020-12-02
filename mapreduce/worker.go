package mapreduce

import (
	"log"
	"net"
	"os/exec"
)

type Worker struct {
	Cwd        string
	Executable string
	OutputDir  string
}

func (w *Worker) RequestInput() (string, bool) {
	conn := w.establishConnection()
	w.sendWorkerReady(conn)

	serverResponse := readFromConn(conn)
	if isServerDoneMsg(serverResponse) {
		return "", false
	}

	return serverResponse, true
}

func (w *Worker) establishConnection() net.Conn {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		log.Fatal(MapReduceError{errCouldNotConnect, err.Error()})
	}
	return conn
}

func isServerDoneMsg(serverResponse string) bool {
	return stringToMessageMap[serverResponse] == ServerDone
}

func (w *Worker) AlertServerOfProgress(info string) {
	conn := w.establishConnection()
	w.sendJobAlert(conn, info)
}

func (w *Worker) ProcessInput(inputFilePath string, outputFilePath string) {
	fullExecutable := w.getFullExecutablePath()
	command := w.buildWorkerCommand(fullExecutable, inputFilePath, outputFilePath)
	command.Run()
}

func (w *Worker) getFullExecutablePath() string {
	return w.Cwd + "/" + w.Executable
}

func (w *Worker) buildWorkerCommand(executable string, inputFilePath string, outputFilePath string) *exec.Cmd {
	commandString := executable + " " + inputFilePath + " " + outputFilePath
	return exec.Command(commandString)
}
