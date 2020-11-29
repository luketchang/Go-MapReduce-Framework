package main

import (
	"os"

	"github.com/mapreduce/server"
)

func main() {
	args := os.Args[1:]
	s := server.NewServer(args)
	s.Run(args)
}
