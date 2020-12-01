package mapper

import (
	"github.com/mapreduce/mapreduce"
)

type Mapper struct {
	mapreduce.Worker
}

func (m *Mapper) startMappingFiles() {
	for {
		inputFilePath, serverDone := m.RequestInput()
		if serverDone {
			break
		}

		m.AlertServerOfProgress("About to process \"" + inputFilePath + "\".")
		m.ProcessInput(inputFilePath)
	}
}
