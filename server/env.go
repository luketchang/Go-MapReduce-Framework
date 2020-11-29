package server

import (
	"log"
	"os"
)

const (
	ServerAddress string = "127.0.0.1:8000"
)

func GetHost() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(errEnvConfig)
	}
	return name
}

func GetCwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(errEnvConfig)
	}
	return path
}
