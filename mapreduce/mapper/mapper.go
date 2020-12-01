package mapper

import (
	"fmt"

	"github.com/mapreduce/mapreduce"
)

type Mapper struct {
	mapreduce.Worker
}

func (m *Mapper) startMappingFiles() {
	for {
		input, serverDone := m.RequestInput()
		if serverDone {
			break
		}
		fmt.Println("Mapper started...")
	}
}
