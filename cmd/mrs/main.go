package main

import (
	"os"

	"github.com/mapreduce/mapreduce"
)

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}

	executable := os.Args[1]
	intermediateDir := os.Args[2]

	sorter := mapreduce.Sorter{
		mapreduce.Worker{
			Executable: executable,
			OutputDir:  intermediateDir,
		},
	}

	sorter.StartSortingFiles()
}
