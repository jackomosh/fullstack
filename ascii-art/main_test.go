package main

import (
	"strings"
	"testing"
)

func TestBannerLoading(t *testing.T) {
	lines, err := getBanner("standard.txt")
	if err != nil {
		t.Fatalf("Could not load standard.txt: %v", err)
	}

	// We check for >= 855 because some files have a trailing newline at the end of the file.
	if len(lines) < 855 {
		t.Errorf("Banner file should have at least 855 lines, but got %d", len(lines))
	}
}

func TestCalculationLogic(t *testing.T) {
	tests := []struct {
		char     rune
		row      int
		expected int
	}{
		{' ', 1, 1},
		{' ', 8, 8},
		{'!', 1, 10},
		{'!', 8, 17},
	}

	for _, tt := range tests {
		result := int(tt.char-32)*9 + tt.row
		if result != tt.expected {
			t.Errorf("Math for %q row %d failed: expected %d, got %d", tt.char, tt.row, tt.expected, result)
		}
	}
}

func TestNewlineSplitting(t *testing.T) {
	input := "Hello\\n\\nThere"
	normalized := strings.ReplaceAll(input, "\\n", "\n")
	parts := strings.Split(normalized, "\n")

	if len(parts) != 3 {
		t.Errorf("Expected 3 segments, got %d", len(parts))
	}

	if parts[1] != "" {
		t.Errorf("Middle segment should be empty, got %q", parts[1])
	}
}