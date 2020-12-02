package mapreduce

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func (s *Server) initializeFromConfigFile(configFilePath string) {
	file := OpenFile(configFilePath)
	defer file.Close()
	s.parseConfigFile(file)
}

func (s *Server) parseConfigFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		key := strings.Split(line, " ")[0]
		if !Contains(ConfigFileKeys, key) {
			log.Fatal(MapReduceError{errBadConfigFile, "non-existent config key"})
		}

		value := strings.Split(line, " ")[1]
		s.applyConfigLineToServer(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(MapReduceError{errScanningFile, err.Error()})
	}
}

func (s *Server) applyConfigLineToServer(key string, value string) {
	//TODO: check string values are valid and enforce range for parseInt
	if key == "mapper" {
		s.mapperExecutable = value
	} else if key == "reducer" {
		s.reducerExecutable = value
	} else if key == "num-mappers" {
		s.numMappers = parseInt(value)
	} else if key == "num-reducers" {
		s.numReducers = parseInt(value)
	} else if key == "input-path" {
		s.inputPath = value
	} else if key == "intermediate-path" {
		s.intermediatePath = value
	} else if key == "output-path" {
		s.outputPath = value
	}
}

func parseInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(MapReduceError{errBadConfigFile, err.Error()})
	}
	return num
}
