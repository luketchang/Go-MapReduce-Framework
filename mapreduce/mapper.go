package mapreduce

import (
	"fmt"
	"log"
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

		m.AlertServerOfProgress("About to process \"" + inputFilePath + "\".")
		m.ProcessInput(inputFilePath, m.OutputDir)
	}
}
