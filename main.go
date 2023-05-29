package main

import (
	"log"
	"os"
	"time"
)

// ProcessMonitor is responsible for monitoring processes and alerting on high CPU usage.
type ProcessMonitor struct {
	cmdExecutor CommandExporter
	Threshold   float64
	Processor   ProcessDataProcessor
	Alerter     Alerter
}

// NewProcessMonitor creates a new instance of ProcessMonitor.
func NewProcessMonitor(cmdExecutor CommandExporter, processor ProcessDataProcessor, alerter Alerter, threshold float64) *ProcessMonitor {
	return &ProcessMonitor{
		cmdExecutor: cmdExecutor,
		Processor:   processor,
		Alerter:     alerter,
		Threshold:   threshold,
	}
}

// Start begins the process monitoring.
func (pm *ProcessMonitor) Start() {
	for {
		if err := pm.Monitor(); err != nil {
			log.Fatalf("Monitoring failed: %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}

// Monitor executes the monitoring process.
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

// AlertOnHighUsage sends alerts for processes with high CPU usage.
func (pm *ProcessMonitor) AlertOnHighUsage(processMetrics []ProcessMetric) {
	for _, metric := range processMetrics {
		if metric.CPU > pm.Threshold {
			pm.Alerter.Alert(metric)
		}
	}
}

// main initializes the necessary components and starts the process monitor.
func main() {
	cmdExecutor := &TopCommandExeporter{}
	processor := ProcessDataProcessor{}
	alerter := &EmailAlerter{}
	threshold := 10.0

	monitor := NewProcessMonitor(cmdExecutor, processor, alerter, threshold)
	monitor.Start()
}
