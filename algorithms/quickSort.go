package algorithms

func QuickSort(numbers []int, low int, high int) {

	if low < high {
		pivot := partition(numbers, low, high)

		QuickSort(numbers, low, pivot-1)
		QuickSort(numbers, pivot+1, high)
	}
}

func partition(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]
	for j := low; j <= high-1; j++ {
		if numbers[j] < pivot {
			i++
			swap(numbers, i, j)
		}
	}

	swap(numbers, i+1, high)

	return i + 1
}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
