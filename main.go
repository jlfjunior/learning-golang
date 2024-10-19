package main

import (
	"fmt"

	s "github.com/jlfjunior/learning-golang/problems"
)

func main() {
	number := 10
	fmt.Printf("Fibonacci of %v: %v", number, s.MemorizedFibonacci((number)))
}
