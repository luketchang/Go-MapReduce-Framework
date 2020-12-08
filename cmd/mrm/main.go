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
	outputDir := os.Args[2]

	mapper := mapreduce.Mapper{
		mapreduce.Worker{
			Executable: executable,
			OutputDir:  outputDir,
		},
	}

	mapper.StartMappingFiles()
}
