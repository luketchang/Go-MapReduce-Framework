package server

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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
