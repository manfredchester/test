package main

import (
	"testing"
)

func top(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}
	return total
}
func MSP() {
	defer func() {
		top(10)
	}()
	top(100)
}
func CMP() {
	top(100)
	top(10)
}
func BenchmarkConnext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MSP()
	}
}
func BenchmarkPanCloud(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CMP()
	}
}
