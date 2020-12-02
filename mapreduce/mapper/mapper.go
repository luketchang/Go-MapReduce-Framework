package mapper

import (
	"github.com/mapreduce/mapreduce"
)

type Mapper struct {
	mapreduce.Worker
}

func (m *Mapper) StartMappingFiles() {
	for {
		inputFilePath, serverDone := m.RequestInput()
		if serverDone {
			break
		}

		m.AlertServerOfProgress("About to process \"" + inputFilePath + "\".")
		outputFilePath := mapreduce.ChangeExtension(inputFilePath, "mapped")
		m.ProcessInput(inputFilePath, outputFilePath)
	}
}
