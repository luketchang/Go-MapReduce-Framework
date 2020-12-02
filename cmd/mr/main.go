package mr

import (
	"os"

	"github.com/mapreduce/mapreduce"
)

func main() {
	args := os.Args[1:]
	server := mapreduce.NewServer(args)
	server.Run()
}
