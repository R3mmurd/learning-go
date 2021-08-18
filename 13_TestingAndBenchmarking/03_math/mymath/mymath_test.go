package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type TestCase struct {
		values []int
		result float64
	}

	tests := []TestCase{
		TestCase{[]int{1, 4, 6, 8, 100}, 6},
		TestCase{[]int{0, 8, 10, 1000}, 9},
		TestCase{[]int{9000, 4, 10, 8, 6, 12}, 9},
		TestCase{[]int{123, 744, 140, 200}, 170},
	}

	for _, test := range tests {
		result := CenteredAvg(test.values)

		if result != test.result {
			t.Error("got", result, "want", test.result)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}
