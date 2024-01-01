func largestAltitude(gain []int) int {
	
    var highest int
	var sum int

	for _, v := range gain{
		sum += v
		if sum > highest{
			highest = sum
		}
	}

    return highest
    
}