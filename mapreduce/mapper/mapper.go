package mapper

import (
	"fmt"
	"log"

	"github.com/mapreduce/mapreduce"
)

type Mapper struct {
	mapreduce.Worker
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
		outputFilePath := mapreduce.ChangeExtension(inputFilePath, "mapped")
		m.ProcessInput(inputFilePath, outputFilePath)
	}
}
