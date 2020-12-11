package mapreduce

import (
	"log"
	"net"
	"os/exec"
)

type Worker struct {
	Executable string
	OutputDir  string
}

func (w *Worker) RequestInput() (string, bool) {
	conn := w.establishConnection()
	w.sendWorkerReady(conn)

	serverResponse := readFromConn(conn)
	if isServerDoneMsg(serverResponse) {
		return "", true
	}

	return serverResponse, false
}

func (w *Worker) establishConnection() net.Conn {
	conn, err := net.Dial("tcp", ExternalServerAddress)
	if err != nil {
		log.Fatal(MapReduceError{errCouldNotConnect, err})
	}

	log.Println("Worker connected to server!")
	return conn
}

func isServerDoneMsg(serverResponse string) bool {
	return stringToMessageMap[serverResponse] == ServerDone
}

func (w *Worker) AlertServerOfProgress(info string) {
	conn := w.establishConnection()
	w.sendJobProgressAlert(conn, info)
}

func (w *Worker) NotifyServerOfJobStatus(fileName string, err error) {
	conn := w.establishConnection()
	if err == nil {
		w.sendJobSucceeded(conn, fileName)
	} else {
		w.sendJobFailed(conn, fileName)
	}
}

func (w *Worker) ProcessInput(inputFilePath string, outputFilePath string) error {
	command := w.buildWorkerCommand(w.Executable, inputFilePath, outputFilePath)
	err := command.Start()
	if err != nil {
		log.Fatal(MapReduceError{errExecutingCmd, err})
	}
	err = command.Wait()
	return err
}

func (w *Worker) buildWorkerCommand(executable string, inputFilePath string, outputFilePath string) *exec.Cmd {
	return exec.Command("sudo", executable, inputFilePath, outputFilePath)
}
