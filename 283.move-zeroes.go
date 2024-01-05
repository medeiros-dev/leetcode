func moveZeroes(nums []int) {

	size := len(nums)

	if size == 1 {
		return
	}

	i := 0
	j := 1

	for j < size && i < size {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[i] == 0 {
			j++
		} else {
			i++
			j++
		}

	}
}