package lc0651

func maxA(n int) int {
	dp := make([]int, n+1)

	for hi := 0; hi <= n; hi++ {
		dp[hi] = hi

		for lo := 1; lo <= hi-3; lo++ {
			curMost := dp[lo] * (hi - lo - 1)
			if curMost > dp[hi] {
				dp[hi] = curMost
			}
		}
	}

	return dp[n]
}
