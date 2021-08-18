package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	hy := 7
	dy := Years(hy)

	if dy != 49 {
		t.Error("Got", dy, "expected", 49)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}
func TestYearsTwo(t *testing.T) {
	hy := 7
	dy := YearsTwo(hy)

	if dy != 49 {
		t.Error("Got", dy, "expected", 49)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
