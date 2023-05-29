package main

import (
	"bufio"
	"log"
	"os"
)

// ReadLinesFromFile reads lines from a file and returns them as a slice of strings
func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
