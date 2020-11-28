package server

import "os"

func GetHost() string {
	name, err := os.Hostname()
	check(err)
	return name
}

func GetCwd() string {
	path, err := os.Getwd()
	check(err)
	return path
}
