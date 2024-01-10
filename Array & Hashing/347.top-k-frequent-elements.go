func topKFrequent(nums []int, k int) []int {
	numbers := make(map[int]int)

	for _, number := range nums {
		numbers[number]++
	}

	mostRepeateds := make([]int, k)
	for i := 0; i < k; i++ {
		mostRepeated := []int{0, 0}
		for number, repeats := range numbers {
			if repeats > mostRepeated[1] {
				mostRepeated[0] = number
				mostRepeated[1] = repeats
			}
		}
		mostRepeateds[i] = mostRepeated[0]
		numbers[mostRepeated[0]] = 0
	}

	return mostRepeateds
}