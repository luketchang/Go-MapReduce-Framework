package mapreduce

import (
	"fmt"
	"log"
)

type Sorter struct {
	Worker
}

func (st *Sorter) StartSortingFiles() {
	for {
		fmt.Println("Sorter worker started...")
		inputFilePath, serverDone := st.RequestInput()
		log.Println(inputFilePath, serverDone)
		if serverDone {
			break
		}

		st.AlertServerOfProgress("About to sort \"" + inputFilePath + "\".")
		st.ProcessInput(inputFilePath)
	}
}
