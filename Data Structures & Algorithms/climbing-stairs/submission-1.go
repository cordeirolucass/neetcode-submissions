func climbStairs(n int) int {
	memo := make([]int, n+1)

	var calculatePossibilities func(step int) int

	calculatePossibilities = func(step int) int {
		if step == 1 { return 1 }
		if step == 2 { return 2 }

		if memo[step] != 0 { return memo[step]}

		result := calculatePossibilities(step-1) + calculatePossibilities(step-2)
		memo[step] = result
		return result
	}
	return calculatePossibilities(n)
}
