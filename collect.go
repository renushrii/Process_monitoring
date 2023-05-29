package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

// CommandExecutor executes commands and collects output.
type CommandExporter interface {
	RunCommand(outputFile string) error
}

// TopCommandExecutor executes the top command and collects the output.
type TopCommandExporter struct{}

// RunCommand executes the top command and writes the output to a file.
func (t *TopCommandExporter) RunCommand() ([]string, error) {
	buffer := bufio.NewScanner(os.Stdout)

	cmd := exec.Command("top", "-b", "-n", "1")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to execute command: %s\n", err)
		return nil, err
	}

	var lines []string
	for buffer.Scan() {
		lines = append(lines, buffer.Text())
	}

	return lines, nil
}
