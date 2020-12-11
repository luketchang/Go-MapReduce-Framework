package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	FailedToOpenCode int = 2
)

func main() {
	time.Sleep(2 * time.Second) //alternate machines
	if len(os.Args[1:]) != 2 {
		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	reduceSortedFile(inputFilePath, outputFilePath)
}

func reduceSortedFile(inputFilePath string, outputFilePath string) {
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
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		wordEndIndex := strings.Index(line, " ")

		word := line[:wordEndIndex]
		counts := line[wordEndIndex+1:]
		countsList := strings.Split(counts, " ")

		totalCount := 0
		for _, num := range countsList {
			inc, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			totalCount += inc
		}

		fmt.Fprintln(outputFile, getOutputPair(word, totalCount))
	}
}

func getOutputPair(word string, count int) string {
	return word + " " + strconv.Itoa(count)
}
