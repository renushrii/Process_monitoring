package main

import (
	"log"
	"strconv"
	"strings"
)

type ProcessMetric struct {
	PID     int
	Command string
	CPU     float64
	Memory  float64
}

func NewProcessMetric(pID int, cpu float64, command string) ProcessMetric {
	return ProcessMetric{
		PID:     pID,
		CPU:     cpu,
		Command: command,
	}
}

type ProcessDataProcessor struct {
	Threshold float64
}

const ProcessDataStartingLineIndex = 7

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
