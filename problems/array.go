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
