package main

import (
	"os"
	"strconv"

	"github.com/mapreduce/mapreduce"
)

func main() {
	if len(os.Args[1:]) != 3 {
		os.Exit(mapreduce.WrongNumArgs)
	}

	executable := os.Args[1]
	intermediateDir := os.Args[2]
	numHashCodes, err := strconv.Atoi(os.Args[3])
	if err != nil {
		os.Exit(mapreduce.InvalidArgsType)
	}

	mapper := mapreduce.Mapper{
		mapreduce.Worker{
			Executable: executable,
			OutputDir:  intermediateDir,
		},
		numHashCodes,
	}

	mapper.StartMappingFiles()
}
