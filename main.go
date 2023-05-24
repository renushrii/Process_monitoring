package main

import (
	"log"
	"os"
)

const threshold = 10.0

func main() {
	err := runMonitoring()
	if err != nil {
		log.Fatalf("Monitoring failed: %v", err)
	}
}

func runMonitoring() error {
	outputFile := os.Getenv("output.txt")

	if err := runTopCommand(outputFile); err != nil {
		log.Fatalf("Failed to run top command: %v", err)
	}

	lines, err := readLinesFromFile(outputFile)
	if err != nil {
		log.Fatalf("Failed to read lines from file: %v", err)
	}

	dataList, err := processData(lines)
	if err != nil {
		log.Fatalf("Failed to read lines from file: %v", err)
	}

	alertDataList := getAlertDataList(dataList)
	err = sendMails(alertDataList)
	if err != nil {
		log.Fatalf("Failed to send mail: %v", err)
	}

	return nil
}
