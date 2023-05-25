package main

import (
	"log"
	"os"
)

type ProcessMonitor struct {
	cmdExecutor CommandExecutor
	threshold   float64
}

func main() {
	cmdExecutor := TopCommandExecutor{}
	threshold := 10.0

	monitor := ProcessMonitor{
		cmdExecutor: cmdExecutor,
		threshold:   threshold,
	}

	err := monitor.Monitor()
	if err != nil {
		log.Fatalf("Monitoring failed: %v", err)
	}
}

func (pm ProcessMonitor) Monitor() error {
	outputFile := os.Getenv("OUTPUT_FILE")

	if err := pm.cmdExecutor.RunCommand(outputFile); err != nil {
		log.Fatalf("Failed to run top command: %v", err)
	}

	lines, err := pm.cmdExecutor.readLinesFromFile(outputFile)
	if err != nil {
		log.Fatalf("Failed to read lines from file: %v", err)
	}

	dataList, err := processData(lines)
	if err != nil {
		log.Fatalf("Failed to read lines from file: %v", err)
	}

	pm.AlertOnHighUsage(dataList)

	return nil
}

func (pm ProcessMonitor) AlertOnHighUsage(dataList []Data) {
	alertDataList := make([]Data, 0)

	for _, data := range dataList {
		if data.isCPUUsageOverThreshold(pm.threshold) {
			alertDataList = append(alertDataList, data)
		}
	}

	err := sendMails(dataList)
	if err != nil {
		log.Fatalf("Failed to send mail: %v", err)
	}
}
