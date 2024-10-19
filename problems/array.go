package problems

func TwoSum(numbers []int, target int) []int {
	visited := make(map[int]int)

	for i, value := range numbers {
		diff := target - value
		x, ok := visited[diff]
		if ok {
			return []int{x, i}
		}
		visited[value] = i
	}

	return []int{}
}

func CustomEvenOddsSort(numbers []int) {
	var quickSort func(numbers []int, reversed bool)
	organize := func(numbers []int) int {
		i := -1
		for j := 0; j < len(numbers); j++ {
			if numbers[j]%2 == 0 {
				i++
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
		return i
	}

	partition := func(numbers []int, reversed bool) int {
		i := -1
		pivot := numbers[len(numbers)-1]
		if reversed {
			for j := 0; j < len(numbers); j++ {
				if numbers[j] >= pivot {
					i++
					numbers[i], numbers[j] = numbers[j], numbers[i]
				}
			}
		} else {
			for j := 0; j < len(numbers); j++ {
				if numbers[j] <= pivot {
					i++
					numbers[i], numbers[j] = numbers[j], numbers[i]
				}
			}
		}

		return i
	}

	quickSort = func(numbers []int, reversed bool) {

		if len(numbers) <= 1 {
			return
		}

		mid := partition(numbers, reversed)

		quickSort(numbers[:mid], reversed)   //Left Side
		quickSort(numbers[mid+1:], reversed) //Right Side
	}

	mid := organize(numbers)

	quickSort(numbers[:mid+1], false)
	quickSort(numbers[mid+1:], true)
}
