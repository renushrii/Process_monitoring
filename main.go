package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ProcessMonitor is responsible for monitoring processes and alerting on high CPU usage.
type ProcessMonitor struct {
	cmdExecutor CommandExporter
	Threshold   float64
	Processor   ProcessDataProcessor
	Alerter     Alerter
	Metrics     *Metrics
}

// NewProcessMonitor creates a new instance of ProcessMonitor.
func NewProcessMonitor(cmdExecutor CommandExporter, processor ProcessDataProcessor, alerter Alerter, threshold float64) *ProcessMonitor {
	return &ProcessMonitor{
		cmdExecutor: cmdExecutor,
		Processor:   processor,
		Alerter:     alerter,
		Threshold:   threshold,
		Metrics:     metrics,
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
	pm.Metrics.UpdateMetrics(processMetrics)

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

// Metrics holds the Prometheus metrics.
type Metrics struct {
	cpuUsage *prometheus.GaugeVec
}

// Initialize Prometheus metrics.
func InitMetrics() *Metrics {
	cpuUsage := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_cpu_usage",
			Help: "Current CPU usage of processes",
		},
		[]string{"pid", "command"},
	)
	prometheus.MustRegister(cpuUsage)

	return &Metrics{
		cpuUsage: cpuUsage,
	}
}

// UpdateMetrics updates the Prometheus metrics based on the process data.
func (m *Metrics) UpdateMetrics(processMetrics []ProcessMetric) {
	for _, metric := range processMetrics {
		m.cpuUsage.WithLabelValues(strconv.Itoa(metric.PID), metric.Command).Set(metric.CPU)
	}
}

// main initializes the necessary components and starts the process monitor.
func main() {
	cmdExecutor := &TopCommandExporter{}
	processor := ProcessDataProcessor{}
	alerter := &EmailAlerter{}
	threshold := 10.0

	metrics := InitMetrics() // Initialize Prometheus metrics

	monitor := NewProcessMonitor(cmdExecutor, processor, alerter, threshold, metrics)
	go monitor.Start()

	// Create HTTP server for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
