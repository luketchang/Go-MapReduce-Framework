package mapper

import (
	"os"

	"github.com/mapreduce/mapreduce"
)

func main() bool {
	if !hasTwoArgs() {
		return false
	}

	executable := os.Args[1]
	outputPath := os.Args[2]

	mapper := Mapper{
		mapreduce.Worker{
			Executable: executable,
			OutputPath: outputPath,
		},
	}

	mapper.startMappingFiles()
	return true
}

func hasTwoArgs() bool {
	return len(os.Args[1:]) != 2
}
