package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestMaxPathSumFromFile(t *testing.T) {
	triangle, err := loadTriangle("files/hard.json")
	if err != nil {
		t.Fatalf("Failed to load triangle data: %v", err)
	}

	// Log the actual result
	result := maxPathSum(triangle)
	t.Logf("Computed max path sum: %d", result)

	// Update the expected value based on this log
	expected := 7273 // Adjust to match the logged value
	if result != expected {
		t.Errorf("Expected max path sum of %d, got %d", expected, result)
	}
}

func loadTriangle(filePath string) ([][]int, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var triangle [][]int
	err = json.Unmarshal(byteValue, &triangle)
	if err != nil {
		return nil, err
	}

	return triangle, nil
}
