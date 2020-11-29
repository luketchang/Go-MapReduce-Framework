package server

import (
	"log"
	"os"
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
