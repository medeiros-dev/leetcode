func topKFrequent(nums []int, k int) []int {

	numsSize := len(nums)

	if numsSize == 1 {
		return []int{nums[0]}
	}

	numbers := make(map[int]int)
	for _, number := range nums {
		numbers[number]++
	}

	mostFrequent := make([][]int, numsSize+1)
	for number, repeats := range numbers {
		mostFrequent[repeats] = append(mostFrequent[repeats], number)
	}

	var result []int
	for i := numsSize; k > len(result); i-- {
		result = append(result, mostFrequent[i]...)
	}

	return result
}