package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const FailedToOpenCode int = 2

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func main() {
	time.Sleep(2 * time.Second) //alternate machines
	if len(os.Args[1:]) != 2 {
		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	parseFile(inputFilePath, outputFilePath)
}

func parseFile(inputFilePath string, outputFilePath string) {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		os.Exit(FailedToOpenCode)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(FailedToOpenCode)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if isAlphaNumeric(word) {
			word = strings.ToLower(word)
			pair := getOutputPair(word)
			fmt.Fprintln(outputFile, pair)
		}
	}
}

func getOutputPair(word string) string {
	return word + " 1"
}
