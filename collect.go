package main

import (
	"log"
	"os"
	"os/exec"
)

type CommandExecutor interface {
	RunCommand(outputFile string) error
}

type TopCommandExecutor struct{}

func (t *TopCommandExecutor) RunCommand(outputFile string) error {
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
