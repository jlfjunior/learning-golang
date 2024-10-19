package main

import (
	"fmt"

	s "github.com/jlfjunior/learning-golang/problems"
)

func main() {

	numbers := []int{2, 7, 11, 15}

	fmt.Println(s.TwoSum(numbers, 9))
}
