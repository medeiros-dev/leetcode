func increasingTriplet(nums []int) bool {
	smallest := math.MaxInt
	greatest := math.MaxInt
	for _, n := range nums {
		if n <= smallest {
			smallest = n
		} else if n <= greatest {
			greatest = n
		} else {
			return true
		}
	}

	return false
}
