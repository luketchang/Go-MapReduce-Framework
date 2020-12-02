package main

import (
	"os"

	"github.com/mapreduce/mapreduce"
	"github.com/mapreduce/mapreduce/mapper"
)

func main() {
	if !hasTwoArgs() {
		return
	}

	executable := os.Args[1]
	outputDir := os.Args[2]

	mapper := mapper.Mapper{
		mapreduce.Worker{
			Executable: executable,
			OutputDir:  outputDir,
		},
	}

	mapper.StartMappingFiles()
}

func hasTwoArgs() bool {
	return len(os.Args[1:]) == 2
}
