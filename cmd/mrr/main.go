package main

import (
	"os"

	"github.com/mapreduce/mapreduce"
)

func main() {
	if len(os.Args[1:]) != 2 {
		os.Exit(mapreduce.WrongNumArgs)
	}

	executable := os.Args[1]
	outputDir := os.Args[2]

	reducer := mapreduce.Reducer{
		mapreduce.Worker{
			Executable: executable,
			OutputDir:  outputDir,
		},
	}

	reducer.StartReducingFiles()
}
