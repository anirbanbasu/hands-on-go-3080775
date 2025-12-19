// testing/basics/begin/main_test.go
package main

import "testing"

// write a test for sum
func TestSum(t *testing.T) {
	got := sum(1, 2, 3, 4, 5)
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	// setup code here
	println("Setting up tests...")

	// run tests
	m.Run()

	// teardown code here
	println("Tearing down tests...")
}
