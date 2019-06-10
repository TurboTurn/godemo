package gotest

import "testing"

// go test -file webbench_test.go -test.bench=".*"
func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
