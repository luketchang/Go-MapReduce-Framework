package mapreduce

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func (s *Server) initializeFromConfigFile(configFilePath string) {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(MapReduceError{errOpeningFile, configFilePath})
	}
	defer file.Close()

	s.parseConfigFile(file)
}

func (s *Server) parseConfigFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		key := strings.Split(line, " ")[0]
		if !Contains(key, ConfigFileKeys) {
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
		s.inputDir = value
		if string(value[len(value)-1]) != "/" {
			s.inputDir = s.inputDir + "/"
		}
	} else if key == "intermediate-path" {
		s.intermediateDir = value
		if string(value[len(value)-1]) != "/" {
			s.intermediateDir = s.intermediateDir + "/"
		}
	} else if key == "output-path" {
		s.outputDir = value
		if string(value[len(value)-1]) != "/" {
			s.outputDir = s.outputDir + "/"
		}
	}
}

func parseInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(MapReduceError{errBadConfigFile, err.Error()})
	}
	return num
}
