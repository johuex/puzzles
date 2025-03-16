package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	climbs := make([]int, n)
	climbs[0] = 1
	climbs[1] = 2
	for i := 2; i < n; i++ {
		climbs[i] = climbs[i-1] + climbs[i-2]
	}
	return climbs[n-1]
}
