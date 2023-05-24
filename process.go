package main

import (
	"log"
	"strconv"
	"strings"
)

func processData(lines []string) ([]Data, error) {
	dataList := make([]Data, 0)

	for i := 7; i < len(lines); i++ {
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

		data := Data{
			pId:     pID,
			cpu:     cpu,
			command: name,
		}

		dataList = append(dataList, data)
	}
	return dataList, nil
}

func getAlertDataList(dataList []Data) []Data {
	alertDataList := make([]Data, 0)

	for _, data := range dataList {
		if data.isCPUUsageOverThreshold() {
			alertDataList = append(alertDataList, data)
		}
	}

	return alertDataList
}
