// testing/benchmarks/begin/main_test.go
package main

import "testing"

// write a benchmark for sum
func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum(numbers...)
	}
}

// write a benchmark for sumAny
func BenchmarkSumAny(b *testing.B) {
	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i
	}
	b.ResetTimer()
	interfaces := make([]interface{}, len(numbers))
	for i, v := range numbers {
		interfaces[i] = v
	}
	for i := 0; i < b.N; i++ {
		sumAny(interfaces)
	}
}
