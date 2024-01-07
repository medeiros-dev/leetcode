func productExceptSelf(nums []int) []int {

	result := []int{1}
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			result = append(result, nums[i-1])
			continue
		}
		product := result[i-1] * nums[i-1]
		result = append(result, product)
	}

	product := 1
	for i := len(nums) - 2; i >= 0; i-- {
		product *= nums[i+1]
		result[i] *= product
	}

	return result

}