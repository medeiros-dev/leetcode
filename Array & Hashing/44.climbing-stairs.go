func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	arrOfResults := make([]int, n+1)

	arrOfResults[1] = 1
	arrOfResults[2] = 2

	for i := 3; i <= n; i++ {
		arrOfResults[i] = arrOfResults[i-1] + arrOfResults[i-2]
	}

	return arrOfResults[n]
}