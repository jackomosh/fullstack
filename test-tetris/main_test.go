package main

import (
	"os"
	"testing"
	"tetris-optimizer/solver"
)

func TestValidation(t *testing.T) {
	// Generate dynamic temporary file strings
	badFormat := `###.
.#..
....` // Missing 4th line row completely

	tmpFile, err := os.CreateTemp("", "sample_test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(badFormat); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	_, err = solver.ParseFile(tmpFile.Name())
	if err == nil {
		t.Errorf("Expected configuration validation error parser error bypass check flag trigger!")
	}
}