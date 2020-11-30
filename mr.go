package main

import (
	"os"

	"github.com/mapreduce/server"
)

func main() {
	args := os.Args[1:]
	server := server.NewServer(args)
	server.Run()
}
