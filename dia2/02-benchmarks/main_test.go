package main

import "testing"

// Para executar o benchmark, use: go test -bench="."
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}
