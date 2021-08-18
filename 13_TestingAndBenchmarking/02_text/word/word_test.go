package word

import (
	"fmt"
	"testing"

	"github.com/learning-go/13_TestingAndBenchmarking/02_text/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("Hello my dear friend, you are my best friend")
	e := map[string]int{
		"Hello":   1,
		"my":      2,
		"dear":    1,
		"friend,": 1,
		"you":     1,
		"are":     1,
		"best":    1,
		"friend":  1,
	}

	for k, v := range e {
		count := m[k]

		if count != v {
			t.Error("go", count, "want", v)
		}
	}

}

func TestCount(t *testing.T) {
	n := Count("Hello my dear friend")
	if n != 4 {
		t.Error("got", n, "want", 4)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello my dear friend"))
	// Output:
	// 4
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
