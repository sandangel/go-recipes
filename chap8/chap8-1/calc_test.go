package chap8_1

import (
	"testing"
	"time"
)

func TestSumInParallel(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

func TestAverageInParallel(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

func BenchmarkSum(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sum(7, 8, 10)
		}
	})
}

func BenchmarkAverage(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Average(7, 8, 10)
		}
	})
}
