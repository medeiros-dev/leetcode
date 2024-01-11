func longestConsecutive(nums []int) int {

	numsSet := make(map[int]bool)

	for _, num := range nums {
		numsSet[num] = true
	}

	longest := 0
	for _, num := range nums {
		if numsSet[num-1] != true {
			counter := 1
			for numsSet[num+counter] != false {
				counter++
			}
			if longest < counter {
				longest = counter
			}
		}
	}

	return longest
}