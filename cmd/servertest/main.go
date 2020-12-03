package main

import (
	"fmt"
	"time"

	"github.com/mapreduce/mapreduce"
)

func main() {
	server := mapreduce.NewServer([]string{
		"./cmd/mr/mr",
		"--mapper",
		"./cmd/mrm/mrm",
		"--reducer",
		"./cmd/mrr/mrr",
		"--config",
		"./mapreduce/files/example.cfg",
	})
	fmt.Println(server, "running...")
	time.Sleep(30 * time.Second)
}
