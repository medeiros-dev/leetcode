func isPalindrome(x int) bool {
    
    if x < 0{
        return false
    }

	slc := []int{}
	for x > 0 {
		slc = append(slc, x%10)
		x = x / 10
	}
		
	for i, j := 0, len(slc) - 1; i < j; i, j = i+1, j-1 {
		if slc[i] != slc[j]{
			return false
		}
	}

	return true
}