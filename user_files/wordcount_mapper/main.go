package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mapreduce/mapreduce"
)

const (
	FailedToOpenCode int = 2
)

func main() {
	time.Sleep(2 * time.Second) //alternate machines
	if len(os.Args[1:]) != 2 {
		return
	}

	inputPath := os.Args[1]
	intermediateDir := os.Args[2]
	parseFile(inputPath, intermediateDir)
}

func parseFile(inputPath string, intermediateDir string) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		os.Exit(FailedToOpenCode)
	}
	defer inputFile.Close()

	outputFilePath := getIntermediateFilePath(inputPath, intermediateDir)
	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(FailedToOpenCode)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		pair := getOutputPair(word)

		fmt.Fprintln(outputFile, pair)
	}
}

func getOutputPair(word string) string {
	return word + " 1"
}

func getIntermediateFilePath(inputPath string, intermediateDir string) string {
	inputFileName := filepath.Base(inputPath)
	intermediateFileName := mapreduce.ChangeExtension(inputFileName, "mapped")
	return intermediateDir + intermediateFileName
}
