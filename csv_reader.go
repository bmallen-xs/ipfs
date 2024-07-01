package main

import (
	"encoding/csv"
	"os"
)

func readCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var cids []string
	for _, record := range records {
		cids = append(cids, record[0])
	}
	return cids, nil
}
