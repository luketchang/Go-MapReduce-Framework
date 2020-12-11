package mapreduce

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Reducer struct {
	Worker
}

func (r *Reducer) StartReducingFiles() {
	ClearDirContents(r.OutputDir)

	for {
		intermediatePattern, serverDone := r.RequestInput()
		if serverDone {
			break
		}

		initialOutputPattern := r.getInitialOutputPattern(intermediatePattern, r.OutputDir)
		r.createInitialSortedOutputFile(intermediatePattern, initialOutputPattern)

		r.AlertServerOfProgress("About to reduce \"" + intermediatePattern + "\".")
		finalOutputPattern := strings.Replace(initialOutputPattern, "*.", "", 1)
		err := r.ProcessInput(initialOutputPattern, finalOutputPattern)

		os.Remove(initialOutputPattern)
		r.NotifyServerOfJobStatus(intermediatePattern, err)
	}
}

func (r *Reducer) getInitialOutputPattern(intermediatePattern string, outputDir string) string {
	intermediateFilePattern := filepath.Base(intermediatePattern)
	outputFileName := ChangeExtension(intermediateFilePattern, "output")
	return outputDir + outputFileName
}

func (r *Reducer) createInitialSortedOutputFile(intermediatePattern string, initialOutputPattern string) {
	command := r.buildSortInitialOutputCommand(intermediatePattern, initialOutputPattern)
	err := command.Start()
	if err != nil {
		log.Fatal(MapReduceError{errExecutingCmd, err.Error()})
	}
	err = command.Wait()
	log.Println("Sort command exited with status:", err)
}

func (r *Reducer) buildSortInitialOutputCommand(intermediatePattern string, initialOutputPattern string) *exec.Cmd {
	commandOpt := fmt.Sprintf(
		"cat %s | sort | python ./mapreduce/group-by-key.py > %s",
		intermediatePattern,
		initialOutputPattern)
	return exec.Command("sudo", "bash", "-c", commandOpt)
}
