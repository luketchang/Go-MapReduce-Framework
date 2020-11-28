package main

import (
	"os"

	"github.com/mapreduce/server"
)

func main() {
	args := os.Args[1:]
	s := &server.MapReduceServer{}
	s.NewServer(args)
}
