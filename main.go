package main

import (
	"fmt"

	s "github.com/jlfjunior/learning-golang/problems"
)

func main() {
	number := 50
	fmt.Printf("Fibonacci of %v: %v", number, s.Fibonacci(number))
}
