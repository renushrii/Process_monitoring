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

type RingBuffer struct {
	data         []ProcessMetric
	size         int
	readIndex    int
	writeIndex   int
	isBufferFull bool
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data:         make([]ProcessMetric, size),
		size:         size,
		readIndex:    0,
		writeIndex:   0,
		isBufferFull: false,
	}
}

func (rb *RingBuffer) Add(processMetric ProcessMetric) {
	rb.data[rb.writeIndex] = processMetric
	rb.writeIndex = (rb.writeIndex + 1) % rb.size

	if rb.writeIndex == rb.readIndex {
		rb.isBufferFull = true
		rb.readIndex = (rb.readIndex + 1) % rb.size
	}
}

func (rb *RingBuffer) GetSlice() []ProcessMetric {
	if !rb.isBufferFull {
		return rb.data[:rb.writeIndex]
	}

	result := make([]ProcessMetric, rb.size)
	copy(result, rb.data[rb.readIndex:])
	copy(result[rb.size-rb.readIndex:], rb.data[:rb.writeIndex])
	return result
}
