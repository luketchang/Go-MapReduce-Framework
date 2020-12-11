package mapreduce

import (
	"log"
	"os"
)

const (
	InternalServerAddress string = ":8000"
	ExternalServerAddress string = "34.94.232.154:8000"
	ServerZone            string = "us-west2-a"
)

func GetHost() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(MapReduceError{errEnvConfig, err})
	}
	return name
}

func GetCwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(MapReduceError{errEnvConfig, err})
	}
	return path
}
