package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// JSONToCSV converts a JSON string to a CSV file
func JSONToCSV(jsonStr string, csvFileName string) error {
	// Unmarshal JSON
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return err
	}

	// Open CSV file for writing
	file, err := os.Create(csvFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to CSV file
	if len(data) > 0 {
		headers := getHeaders(data[0])
		if err := writer.Write(headers); err != nil {
			return err
		}
	}

	// Write data to CSV file
	for _, record := range data {
		var row []string
		for _, header := range getHeaders(record) {
			value := fmt.Sprintf("%v", record[header])
			row = append(row, value)
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ByKey []KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// getHeaders returns the headers of a map
func getHeaders(m map[string]interface{}) []string {
	var headers []string
	for k := range m {
		headers = append(headers, k)
	}
	return headers
}

// JSONStringToValues converts a JSON string to a slice of strings containing JSON values
func JSONStringToValues(jsonStr string) ([]string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}

	// Create a slice to store key-value pairs
	var keyValueSlice []KeyValue

	// Iterate over the map and populate the slice
	for key, value := range data {
		keyValueSlice = append(keyValueSlice, KeyValue{Key: key, Value: value})
	}

	// Sort the slice based on keys
	sort.Sort(ByKey(keyValueSlice))

	var values []string

	for _, keyValue := range keyValueSlice {
		fmt.Println(keyValue.Key)
		values = append(values, keyValue.Value.(string))
	}
	return values, nil
}
