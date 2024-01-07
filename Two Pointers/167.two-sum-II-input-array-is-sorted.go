func twoSum(numbers []int, target int) []int {

	pntOne := 0
	pntTwo := len(numbers) - 1

	for true {
		sum := numbers[pntOne] + numbers[pntTwo]
		if sum > target {
			pntTwo--
		} else if sum < target {
			pntOne++
		} else {
			break
		}
	}

	return []int{pntOne + 1, pntTwo + 1}
}