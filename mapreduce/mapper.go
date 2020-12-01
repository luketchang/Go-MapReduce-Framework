package mapreduce

import "fmt"

type Mapper struct {
	Worker
}

func (m *Mapper) mapFiles() {
	for {
		// input, serverDone := m.requestInput()
		// if serverDone {
		// 	break
		// }
		fmt.Println("Mapper started...")
	}
}
