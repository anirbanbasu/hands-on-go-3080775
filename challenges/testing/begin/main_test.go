// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCounter_Count(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello, World!", 10},
		{"12345", 0},
		{"GoLang", 6},
		{"", 0},
	}

	for _, test := range tests {
		counter := letterCounter{}
		t.Run(test.input, func(t *testing.T) {
			result := counter.count(test.input)
			if result != test.expected {
				t.Errorf("For input %q, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter_Count(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello, World! 123", 3},
		{"No numbers here!", 0},
		{"2024 is the year", 4},
		{"", 0},
	}

	for _, test := range tests {
		counter := numberCounter{}
		t.Run(test.input, func(t *testing.T) {
			result := counter.count(test.input)
			if result != test.expected {
				t.Errorf("For input %q, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCounter_Count(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello, World! @2024", 3},
		{"No symbols here", 0},
		{"#GoLang is awesome!", 2},
		{"", 0},
	}

	for _, test := range tests {
		counter := symbolCounter{}
		t.Run(test.input, func(t *testing.T) {
			result := counter.count(test.input)
			if result != test.expected {
				t.Errorf("For input %q, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}
