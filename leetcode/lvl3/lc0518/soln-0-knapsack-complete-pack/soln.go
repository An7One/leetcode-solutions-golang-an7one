package lc0518

func change(amount int, coins []int) int {
	var dp = make([]int, 1+amount)
	dp[0] = 1

	for _, coin := range coins {
		for faceValue := coin; faceValue <= amount; faceValue++ {
			dp[faceValue] += dp[faceValue-coin]
		}
	}

	return dp[amount]
}
