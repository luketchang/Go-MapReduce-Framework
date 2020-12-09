package mapreduce

import (
	"bufio"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Mapper struct {
	Worker
}

func (m *Mapper) StartMappingFiles() {
	for {
		fmt.Println("Mapper worker started...")
		inputFilePath, serverDone := m.RequestInput()
		log.Println(inputFilePath, serverDone)
		if serverDone {
			break
		}

		m.AlertServerOfProgress("About to map \"" + inputFilePath + "\".")
		intermediateFilePath := m.getIntermediateFilePath(inputFilePath, m.OutputDir)
		m.ProcessInput(inputFilePath, intermediateFilePath)
		m.sortMappedFile(intermediateFilePath)
	}
}

func (m *Mapper) getIntermediateFilePath(inputPath string, intermediateDir string) string {
	inputFileName := filepath.Base(inputPath)
	intermediateFileName := ChangeExtension(inputFileName, "mapped")
	return intermediateDir + intermediateFileName
}

func (m *Mapper) sortMappedFile(mappedFilePath string) {
	mappedFile, err := os.Open(mappedFilePath)
	if err != nil {
		log.Fatal(MapReduceError{errOpeningFile, err.Error()})
	}
	defer mappedFile.Close()
	defer os.Remove(mappedFilePath)

	scanner := bufio.NewScanner(mappedFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		key := m.getKey(line)
		hashVal := m.getHashValue(key)
		outputFileName := m.getBucketedFilePath(mappedFilePath, hashVal)
		outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(MapReduceError{errOpeningFile, err.Error()})
		}

		fmt.Fprintln(outputFile, line)
		outputFile.Close()
	}
}

func (m *Mapper) getKey(line string) string {
	return line[:strings.IndexByte(line, ' ')]
}

func (m *Mapper) getHashValue(word string) int {
	h := sha1.New()
	h.Write([]byte(word))
	hash := int(binary.BigEndian.Uint64(h.Sum(nil)))
	if hash < 0 {
		hash = hash * -1
	}

	return hash % 4
}

func (m *Mapper) getBucketedFilePath(mappedFilePath string, hashVal int) string {
	return mappedFilePath + "." + strconv.FormatInt(int64(hashVal), 10)
}
