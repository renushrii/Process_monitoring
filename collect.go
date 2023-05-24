package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func runTopCommand(outputFile string) error {
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return err
	}
	defer f.Close()

	cmd := exec.Command("top", "-b", "-n", "1")
	cmd.Stdin = os.Stdin
	cmd.Stdout = f

	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to run command: %s\n", err)
	}

	return nil
}

func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
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
