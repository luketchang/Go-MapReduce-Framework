package mapreduce

import (
	"os"
)

func main() {
	args := os.Args[1:]
	server := NewServer(args)
	server.Run()
}
