// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCounter_Count(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"hello-world": {input: "Hello, World!", expected: 10},
		"numeric":     {input: "12345", expected: 0},
		"golang":      {input: "GoLang", expected: 6},
		"empty":       {input: "", expected: 0},
	}

	counter := letterCounter{}
	for key, tc := range tests {
		t.Run(key, func(t *testing.T) {
			result := counter.count(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter_Count(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"hello-world": {"Hello, World! 123", 3},
		"no-numbers":  {"No numbers here!", 0},
		"year":        {"2024 is the year", 4},
		"empty":       {"", 0},
	}

	counter := numberCounter{}
	for key, tc := range tests {
		t.Run(key, func(t *testing.T) {
			result := counter.count(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCounter_Count(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"hello-world": {"Hello, World! @2024", 3},
		"no-symbols":  {"No symbols here", 0},
		"go-lang":     {"#GoLang is awesome!", 2},
		"empty":       {"", 0},
	}

	for key, tc := range tests {
		counter := symbolCounter{}
		t.Run(key, func(t *testing.T) {
			result := counter.count(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}
