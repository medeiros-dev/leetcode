func plusOne(digits []int) []int {

	lenDigits := len(digits)

	for i := lenDigits - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits[0] = 1
	digits = append(digits, 0)

	return digits

}