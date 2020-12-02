package mrm

import (
	"os"

	"github.com/mapreduce/mapreduce"
	"github.com/mapreduce/mapreduce/mapper"
)

func main() bool {
	if !hasTwoArgs() {
		return false
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
	return true
}

func hasTwoArgs() bool {
	return len(os.Args[1:]) != 2
}
