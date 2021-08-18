package main

import (
	"fmt"

	"github.com/learning-go/13_TestingAndBenchmarking/02_text/quote"
	"github.com/learning-go/13_TestingAndBenchmarking/02_text/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
