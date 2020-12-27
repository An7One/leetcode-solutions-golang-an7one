package lc0877

func stoneGame(piles []int) bool {
	dp := make([][]int, len(piles))
	for i := range dp {
		dp[i] = make([]int, len(piles))
	}

	for idx := range piles {
		dp[idx][idx] = piles[idx]
	}

	for lens := 1; lens < len(piles); lens++ {
		for lo := 0; lo < len(piles)-lens; lo++ {
			hi := lo + lens

			if piles[lo]-dp[lo+1][hi] > piles[hi]-dp[lo][hi-1] {
				dp[lo][hi] = piles[lo] - dp[lo+1][hi]
			} else {
				dp[lo][hi] = piles[hi] - dp[lo][hi-1]
			}
		}
	}

	return dp[0][len(piles)-1] > 0
}
