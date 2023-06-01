package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// ProcessMetric represents the metrics of a process.
type ProcessMetric struct {
	PID     int
	Command string
	CPU     float64
	Memory  float64
}

// NewProcessMetric creates a new instance of ProcessMetric.
func NewProcessMetric(pID int, cpu float64, command string) ProcessMetric {
	return ProcessMetric{
		PID:     pID,
		CPU:     cpu,
		Command: command,
	}
}

// ProcessDataProcessor is responsible for processing process data.
type ProcessDataProcessor struct {
	Threshold      float64
	BufferDuration time.Duration
	Buffers        map[int]*RingBuffer
}

// ProcessDataStartingLineIndex is the index of the starting line in the process data
const ProcessDataStartingLineIndex = 7

// ProcessData processes the lines and returns the corresponding process metrics.
func (pdp *ProcessDataProcessor) ProcessData(lines []string) []ProcessMetric {
	processMetrics := make([]ProcessMetric, 0)

	for i := ProcessDataStartingLineIndex; i < len(lines); i++ {
		line := lines[i]
		words := strings.Fields(line)

		pID, err := strconv.Atoi(words[0])
		if err != nil {
			log.Printf("Error converting PID: %s\n", err)
			continue
		}

		cpu, err := strconv.ParseFloat(words[8], 64)
		if err != nil {
			log.Printf("Error converting CPU value: %s\n", err)
			continue
		}

		name := words[11]

		processMetric := NewProcessMetric(pID, cpu, name)

		processMetrics = append(processMetrics, processMetric)
	}
	return processMetrics
}
