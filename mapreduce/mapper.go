package mapreduce

import (
	"fmt"
	"log"
	"path/filepath"
)

type Mapper struct {
	Worker
}

func (m *Mapper) StartMappingFiles() {
	for {
		fmt.Println("Mapper worker started...")
		inputFilePath, serverDone := m.RequestInput()
		log.Println(inputFilePath, serverDone)
		if serverDone {
			break
		}

		m.AlertServerOfProgress("About to map \"" + inputFilePath + "\".")
		intermediateFilePath := getIntermediateFilePath(inputFilePath, m.OutputDir)
		m.ProcessInput(inputFilePath, intermediateFilePath)
	}
}

func getIntermediateFilePath(inputPath string, intermediateDir string) string {
	inputFileName := filepath.Base(inputPath)
	intermediateFileName := ChangeExtension(inputFileName, "mapped")
	return intermediateDir + intermediateFileName
}

func (m *Mapper) sortMappedFiles() {
}
