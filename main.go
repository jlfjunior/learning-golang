package main

import (
	"fmt"

	s "github.com/jlfjunior/learning-golang/problems"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Unsorted: ", numbers)
	s.CustomEvenOddsSort(numbers)
	fmt.Println("Sorted: ", numbers)
}
