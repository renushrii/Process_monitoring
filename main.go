package main

import (
	"log"
	"os"
	"time"
)

type ProcessMonitor struct {
	cmdExecutor CommandExecutor
	threshold   float64
	Processor   ProcessDataProcessor
	Alerter     Alerter
}

func NewProcessMonitor(cmdExecutor CommandExecutor, processor ProcessDataProcessor, alerter Alerter, threshold float64) *ProcessMonitor {
	return &ProcessMonitor{
		cmdExecutor: cmdExecutor,
		Processor:   processor,
		Alerter:     alerter,
		threshold:   threshold,
	}
}

func (pm *ProcessMonitor) Start() {
	for {
		if err := pm.Monitor(); err != nil {
			log.Fatalf("Monitoring failed: %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	cmdExecutor := &TopCommandExecutor{}
	processor := ProcessDataProcessor{}
	alerter := &EmailAlerter{}
	threshold := 10.0

	monitor := NewProcessMonitor(cmdExecutor, processor, alerter, threshold)
	monitor.Start()
}

func (pm *ProcessMonitor) Monitor() error {
	outputFile := os.Getenv("OUTPUT_FILE")

	if err := pm.cmdExecutor.RunCommand(outputFile); err != nil {
		log.Fatalf("Failed to run top command: %v", err)
	}

	lines, err := ReadLinesFromFile(outputFile)
	if err != nil {
		log.Fatalf("Failed to read lines from file: %v", err)
	}

	processMetrics := pm.Processor.ProcessData(lines)
	pm.AlertOnHighUsage(processMetrics)

	return nil
}

func (pm *ProcessMonitor) AlertOnHighUsage(processMetrics []ProcessMetric) {
	for _, metric := range processMetrics {
		if metric.CPU > pm.threshold {
			pm.Alerter.Alert(metric)
		}
	}
}
