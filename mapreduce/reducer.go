package mapreduce

import (
	"log"
	"os/exec"
	"path/filepath"
)

type Reducer struct {
	Worker
}

func (r *Reducer) StartReducingFiles() {
	for {
		intermediatePattern, serverDone := r.RequestInput()
		if serverDone {
			break
		}

		initialOutputPattern := r.getOutputFilePattern(intermediatePattern, r.OutputDir)
		r.createInitialSortedOutputFile(intermediatePattern, initialOutputPattern)

		r.AlertServerOfProgress("About to reduce \"" + intermediatePattern + "\".")
		//TODO: call wordcount reducer executable
	}
}

func (r *Reducer) getOutputFilePattern(intermediatePattern string, outputDir string) string {
	intermediateFilePattern := filepath.Base(intermediatePattern)
	outputFileName := ChangeExtension(intermediateFilePattern, "out")
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
	return exec.Command(
		"cat", intermediatePattern,
		"|", "sort",
		"|", "python3", "./mapreduce/group-by-key.py",
		">", initialOutputPattern)
}
