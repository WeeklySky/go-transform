package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotransform <file.csv>")
		return
	}

	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	if len(records) < 1 {
		fmt.Println("The CSV file is empty or does not have headers.")
		return
	}

	headers := records[0]
	var jsonData []map[string]string

	for _, row := range records[1:] {
		if len(row) != len(headers) {
			fmt.Println("Error: number of fields in a row is different.")
			return
		}

		record := make(map[string]string)
		for i, header := range headers {
			record[header] = row[i]
		}
		jsonData = append(jsonData, record)
	}

	jsonOutput, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error while converting to JSON:", err)
		return
	}

	jsonFileName := strings.TrimSuffix(os.Args[1], ".csv") + ".json"
	jsonFile, err := os.Create(jsonFileName)
	if err != nil {
		fmt.Println("Error while creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonOutput)
	if err != nil {
		fmt.Println("Error while writing data into JSON file:", err)
		return
	}

	fmt.Printf("JSON file created with success: %s", jsonFileName)
}
