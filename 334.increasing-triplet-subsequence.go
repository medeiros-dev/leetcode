func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt
	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}

	return false
}