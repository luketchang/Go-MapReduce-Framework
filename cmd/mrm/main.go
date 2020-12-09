package main

import (
	"os"
	"strconv"

	"github.com/mapreduce/mapreduce"
)

const (
	WrongNumArgs       int = 2
	NumHashCodesNotInt int = 3
)

func main() {
	if len(os.Args[1:]) != 3 {
		os.Exit(WrongNumArgs)
	}

	executable := os.Args[1]
	intermediateDir := os.Args[2]
	numHashCodes, err := strconv.Atoi(os.Args[3])
	if err != nil {
		os.Exit(NumHashCodesNotInt)
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
