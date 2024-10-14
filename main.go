package main

import sort "github.com/jlfjunior/learning-golang/sort"

func main() {
	numbers := []int{9, 0, 3, 7, 5, 2, 8, 4, 6, 1}
	sort.QuickSort(numbers, 0, len(numbers)-1)
}
